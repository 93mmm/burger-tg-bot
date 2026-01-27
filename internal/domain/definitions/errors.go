package definitions

import "github.com/pkg/errors"

var (
	// общая ошибка отсутствия чего-либо.
	ErrNotFound = errors.New("объект не найден")
	// внутренняя ошибка сервиса.
	ErrInternal = errors.New("внутренняя ошибка сервиса")
	// данные запроса не позволяют выполнить операцию
	ErrInvalidArgument = errors.New("данные запроса не позволяют выполнить операцию")
	// отказано в доступе, нет прав
	ErrAccessDenied = errors.New("отказано в доступе")
	// не авторизован
	ErrNotAuthorized = errors.New("не авторизован")
	// внешняя ошибка, не можем получить детальную информацию
	ErrExternalSystem = errors.New("внешняя ошибка")
)
