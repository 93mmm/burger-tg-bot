package logger

import (
	"context"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	errMsgOddNumber    = "Ignored key without a value."
	errMsgNonStringKey = "Ignored key-value pairs with non-string keys."
	errMsgMultiple     = "Multiple errors without a key."
)

type invalidPair struct {
	position   int
	key, value any
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	zap.Any("key", p.key).AddTo(enc)
	zap.Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}

func mergeKvs(ctx context.Context, otherKVs ...any) []any {
	kvsFromContext := getKvsFromContext(ctx)

	// Ранее ключи не были добавлены.
	if len(kvsFromContext) == 0 {
		return sweetenFields(ctx, otherKVs)
	}

	return mergeFields(kvsFromContext, sweetenFields(ctx, otherKVs))
}

func mergeFields(fieldsFromContext, finalCallFields []any) []any {
	merged := make([]any, 0, len(fieldsFromContext)+len(finalCallFields))
	merged = append(merged, fieldsFromContext...)

	for i := 0; i < len(finalCallFields); i++ {
		wasFieldFromContextReplaced := false
		newField, ok := finalCallFields[i].(zap.Field)
		if !ok {
			continue
		}

		for j := range merged {
			fieldFromContext := merged[j].(zap.Field)

			if fieldFromContext.Key == newField.Key {
				merged[j] = newField
				wasFieldFromContextReplaced = true
				break
			}
		}

		if !wasFieldFromContextReplaced {
			merged = append(merged, newField)
		}
	}

	return merged
}

func sweetenFields(ctx context.Context, args []any) []any {
	if len(args) == 0 {
		return nil
	}

	var (
		fields    = make([]any, 0, len(args))
		invalid   invalidPairs
		seenError bool
	)
	logger := getLogger(ctx)
	for i := 0; i < len(args); {
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		if err, ok := args[i].(error); ok {
			if !seenError {
				seenError = true
				fields = append(fields, zap.Error(err))
			} else {
				logger.Error(errMsgMultiple, zap.Error(err))
			}
			i++
			continue
		}

		if i == len(args)-1 {
			logger.Error(errMsgOddNumber, zap.Any("ignored", args[i]))
			break
		}

		// Находим пару ключ-значение.
		key, val := args[i], args[i+1]

		if keyStr, ok := key.(string); !ok {
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}

			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
		i += 2
	}
	if len(invalid) > 0 {
		logger.Error(errMsgNonStringKey, zap.Array("invalid", invalid))
	}
	return fields
}
