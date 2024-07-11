package I18n

// Option isn't added to the main dictionary
// as some invalid translations present
// e.g. DelayThreshold_OnlyDiscrete, missing constant string "discrete"

var I18nDict_ru_RU map[uint16]string = map[uint16]string{
	ACME_FailedToGetCommand:             "Не удалось прочитать команду ACME",
	ACME_FailedToSeek:                   "Недопустимый файл ACME, поскольку операция поиска завершилась неудачей.",
	ACME_StructureErrorNotice:           "Ошибка файловой структуры",
	ACME_UnknownCommand:                 "Неизвестная команда ACME (ошибка файла)",
	Auth_BackendError:                   "Ошибка серверной части",
	Auth_FailedToRequestEntry:           "Не удается запросить вход в службу проката, пожалуйста, проверьте, отключена ли настройка уровня обслуживания проката и правильный ли пароль службы проката.",
	Auth_HelperNotCreated:               "Вспомогательный пользователь еще не создан, пожалуйста, перейдите в центр пользователей, чтобы создать его.",
	Auth_InvalidFBVersion:               "Версия FastBuilder недействительна, пожалуйста, обновите.",
	Auth_InvalidHelperUsername:          "Имя пользователя вспомогательного пользователя неверно, пожалуйста, перейдите в центр пользователя, чтобы настроить его.",
	Auth_InvalidToken:                   "Недействительный токен, пожалуйста, войдите в систему еще раз.",
	Auth_InvalidUser:                    "Недопустимый пользователь, пожалуйста, войдите в систему еще раз.",
	Auth_ServerNotFound:                 "Услуга проката не найдена, пожалуйста, проверьте, открыта ли услуга проката для всех.",
	Auth_UnauthorizedRentalServerNumber: "Соответствующий номер службы проката не был авторизован, пожалуйста, перейдите в центр пользователей для авторизации.",
	Auth_UserCombined:                   "Пользователь был объединен с другой учетной записью, пожалуйста, войдите в систему с новой учетной записью.",
	Auth_FailedToRequestEntry_TryAgain:  "Не удается запросить вход в службу проката, пожалуйста, повторите попытку позже.",
	BDump_Author:                        "автор",
	BDump_EarlyEOFRightWhenOpening:      "Файл не удалось прочитать, поскольку он завершился преждевременно и, возможно, был поврежден.",
	BDump_FailedToGetCmd1:               "Не удалось получить какие-либо параметры cmd [pos:0], возможно, файл поврежден",
	BDump_FailedToGetCmd2:               "Не удалось получить какие-либо параметры cmd [pos1], возможно, файл поврежден",
	BDump_FailedToGetCmd4:               "Не удалось получить какие-либо параметры cmd [pos2], возможно, файл поврежден",
	BDump_FailedToGetCmd6:               "Не удалось получить какие-либо параметры cmd [pos3], возможно, файл поврежден",
	BDump_FailedToGetCmd7_0:             "Не удалось получить какие-либо параметры cmd [pos4], возможно, файл поврежден",
	BDump_FailedToGetCmd7_1:             "Не удалось получить какие-либо параметры cmd [pos5], возможно, файл поврежден",
	BDump_FailedToGetCmd10:              "Не удалось получить какие-либо параметры cmd [pos6], возможно, файл поврежден",
	BDump_FailedToGetCmd11:              "Не удалось получить какие-либо параметры cmd [pos7], возможно, файл поврежден",
	BDump_FailedToGetCmd12:              "Не удалось получить какие-либо параметры cmd [pos8], возможно, файл поврежден",
	BDump_FailedToGetConstructCmd:       "Не удалось получить команду сборки, возможно, файл поврежден",
	BDump_FailedToReadAuthorInfo:        "Не удалось прочитать информацию об авторе, возможно, файл поврежден",
	BDump_FileNotSigned:                 "Файл без подписи",
	BDump_FileSigned:                    "Документ подписан, владелец：%s",
	BDump_NotBDX_Invheader:              "Не является файлом bdx (недопустимый заголовок файла)",
	BDump_NotBDX_Invinnerheader:         "Не является файлом bdx (недопустимый внутренний заголовок файла)",
	BDump_SignedVerifying:               "Файл подписан и проходит проверку...",
	BDump_VerificationFailedFor:         "потому что %v Не удалось проверить подпись файла.",
	BDump_Warn_Reserved:                 "предупреждать：BDump/Import：Используются зарезервированные поля\n",
	CommandNotFound:                     "Эта команда не может быть найдена",
	ConnectionEstablished:               "Успешно подключено к серверу.",
	Crashed_No_Connection:               "Не удавалось установить соединение в течение длительного времени",
	Crashed_OS_Windows:                  "Нажмите ENTER, чтобы выйти из программы.",
	Crashed_StackDump_And_Error:         "Дамп стека показан выше.Ошибка заключается в：",
	Crashed_Tip:                         "Столкнулся с проблемой во время работы FastBuilder Phoenix",
	CurrentDefaultDelayMode:             "Текущий режим задержки по умолчанию",
	CurrentTasks:                        "Список задач：",
	DelayModeSet:                        "Установлен режим задержки",
	DelayModeSet_DelayAuto:              "Значение задержки было автоматически установлено на: %d",
	DelayModeSet_ThresholdAuto:          "Порог задержки был автоматически установлен на: %d",
	DelaySet:                            "Задержка была установлена",
	DelaySetUnavailableUnderNoneMode:    "[delay set] Недоступно в режиме \"none\"",
	DelayThreshold_OnlyDiscrete:         "Порог задержки может быть установлен только в дискретном режиме.",
	DelayThreshold_Set:                  "Порог задержки был установлен на %d",
	ERRORStr:                            "ошибающийся",
	EnterPasswordForFBUC:                "Пожалуйста, введите свой пароль для входа в FastBuilder user Center (не будет отображаться): ",
	Enter_FBUC_Username:                 "Введите свое имя пользователя FastBuilder User Center: ",
	Enter_Rental_Server_Code:            "Пожалуйста, введите номер службы проката: ",
	Enter_Rental_Server_Password:        "Введите пароль службы проката (если он не установлен, нажмите [Enter] напрямую, ввод не будет повторен): ",
	ErrorIgnored:                        "Эта ошибка была проигнорирована",
	Error_MapY_Exceed:                   "При использовании рисования стереокарты значение MapY должно находиться в диапазоне от 20 до 255 (введенное вами значение равно %v)",
	FBUC_LoginFailed:                    "Имя пользователя или пароль пользовательского центра FastBuilder неверны",
	FBUC_Token_ErrOnCreate:              "При создании файла токена произошла ошибка：",
	FBUC_Token_ErrOnGen:                 "Произошла ошибка при создании временного токена",
	FBUC_Token_ErrOnRemove:              "Не удалось удалить файл токена: %v",
	FBUC_Token_ErrOnSave:                "При сохранении токена произошла ошибка：",
	FileCorruptedError:                  "Файл был поврежден",
	Get_Warning:                         "Команда get была заменена командой set. Эта команда будет удалена в будущем. Пожалуйста, перейдите на set.",
	IgnoredStr:                          "Игнорируется",
	InvalidFileError:                    "Недопустимый файл",
	InvalidPosition:                     "Никаких действительных координат получено не было.(Эту ошибку можно проигнорировать)",
	Lang_Config_ErrOnCreate:             "Произошла ошибка при создании файла конфигурации языка：%v",
	Lang_Config_ErrOnSave:               "Произошла ошибка при сохранении языковой конфигурации：%v",
	LanguageName:                        "Русский",
	LanguageUpdated:                     "Языковые предпочтения были обновлены",
	Logout_Done:                         "Вышел из пользовательского центра FastBuilder.",
	Menu_BackButton:                     "< вернуться",
	Menu_Cancel:                         "отменять",
	Menu_CurrentPath:                    "Текущий путь",
	Menu_ExcludeCommandsOption:          "Исключить содержимое командного поля",
	Menu_GetEndPos:                      "Получить координаты конечной точки",
	Menu_GetPos:                         "Получить координаты",
	Menu_InvalidateCommandsOption:       "Аннулирование команды",
	Menu_Quit:                           "Выйдите из программы",
	Menu_StrictModeOption:               "Строгий режим",
	NotAnACMEFile:                       "Предоставленные документы не являются строительными документами ACME",
	Notice_CheckUpdate:                  "Проверяя наличие обновлений, пожалуйста, подождите…",
	Notice_iSH_Location_Service:         "Вы используете эмулятор iSH, и для его работы в фоновом режиме необходимо использовать разрешение на определение местоположения. Кроме этого, никакие данные о местоположении не записываются и не используются. Вы можете отключить его в любое время.",
	Notice_OK:                           "закончить\n",
	Notice_UpdateAvailable:              "Доступна новая версия PhoenixBuilder (%s).\n",
	Notice_UpdateNotice:                 "Пожалуйста, обновите это программное обеспечение.\n",
	Notice_ZLIB_CVE:                     "Ваша версия zlib (%s) содержит подтвержденную критическую уязвимость, мы рекомендуем вам обновить ее, чтобы избежать несчастных случаев",
	Notify_NeedOp:                       "Для правильной работы требуется разрешение OP.",
	Notify_TurnOnCmdFeedBack:            "Вам нужно, чтобы sendcommandfeedback был верным. Мы включили эту опцию для вас. Пожалуйста, отключите ее по мере необходимости после ее использования.",
	Omega_WaitingForOP:                  "Система Omega ожидает разрешения на операцию",
	Omega_Enabled:                       "Система Omega включена!",
	OpPrivilegeNotGrantedForOperation:   "Разрешение на управление роботом не предоставлено, пожалуйста, предоставьте разрешение на управление, прежде чем делать это",
	Original_Important_Contributor:      "Важный вклад в первые дни：LNSSPsd(Ruphane), CAIMEO, CMA2401PT",
	Parsing_UnterminatedEscape:          "Побег не завершен",
	Parsing_UnterminatedQuotedString:    "Заключенная в кавычки часть строки не завершается",
	PositionGot:                         "Получены координаты начальной точки",
	PositionGot_End:                     "Получены координаты конечной точки",
	PositionSet:                         "Координаты были установлены",
	PositionSet_End:                     "Координаты конечной точки были установлены",
	QuitCorrectly:                       "Выходите нормально",
	Sch_FailedToResolve:                 "Не удалось проанализировать файл",
	SelectLanguageOnConsole:             "Пожалуйста, выберите новый язык в консоли",
	ServerCodeTrans:                     "Номер сервера",
	SimpleParser_Int_ParsingFailed:      "Анализатор: Не удалось обработать целочисленные параметры",
	SimpleParser_InvEnum:                "Анализатор: Недопустимое значение перечисления, доступные значения：%s.",
	SimpleParser_Invalid_decider:        "Анализатор: Недопустимый определитель",
	SimpleParser_Too_few_args:           "Анализатор: Слишком мало параметров",
	SkipMCPCheckChallengesLimit:         `В связи с включением опции "skip-mcpc-check-challenges", Функции, используемые в настоящее время, были отключены.`,
	Special_Startup:                     "Включенный язык: Русский\n",
	/* Special SysError Translations, for innocent kids */
	SysError_EACCES:             "В разрешении отказано, пожалуйста, проверьте, разрешен ли программе доступ к соответствующему файлу.",
	SysError_EBUSY:              "Файл занят, пожалуйста, повторите попытку позже.",
	SysError_EINVAL:             "Неверный ввод файла.",
	SysError_EISDIR:             "Входной файл представляет собой каталог, и входные данные недопустимы.",
	SysError_ENOENT:             "Соответствующий файл не существует.",
	SysError_ETXTBSY:            "Файл занят, пожалуйста, повторите попытку позже.",
	SysError_HasTranslation:     "Ошибка в работе с файлом для %s: %s",
	TaskCreated:                 "Задача создана",
	TaskDisplayModeSet:          "Режим отображения статуса задачи был установлен на: %s.",
	TaskFailedToParseCommand:    "Не удалось разобрать команду: %v",
	TaskNotFoundMessage:         "Не удалось найти допустимую задачу на основе предоставленного идентификатора задачи.",
	TaskPausedNotice:            "[миссия %d] - Подвешенный",
	TaskResumedNotice:           "[миссия %d] - восстановленный",
	TaskStateLine:               "ID %d - Командная строка:\"%s\", состояние: %s, Значение задержки: %d, Режим задержки: %s, Порог задержки: %d",
	TaskStoppedNotice:           "[миссия %d] - Остановился",
	TaskTTeIuKoto:               "миссия",
	TaskTotalCount:              "сумма：%d",
	TaskTypeCalculating:         "Вычислительный",
	TaskTypeDied:                "Мертвый",
	TaskTypePaused:              "Подвешенный",
	TaskTypeRunning:             "В действии",
	TaskTypeSpecialTaskBreaking: "Специальная миссия: завершение",
	TaskTypeSwitchedTo:          "Тип создания задачи был переключен на：%s.",
	TaskTypeUnknown:             "неизвестная",
	Task_D_NothingGenerated:     "[миссия %d] Ни одна структура не была успешно сгенерирована.",
	Task_DelaySet:               "[миссия %d] - Задержка была установлена: %d",
	Task_ResumeBuildFrom:        "Возобновите сборку с %v-го квадрата",
	Task_SetDelay_Unavailable:   "[setdelay] Недоступно в режиме без задержки",
	Task_Summary_1:              "[миссия %d] %v квадраты были изменены.",
	Task_Summary_2:              "[миссия %d] Время: %v секунд",
	Task_Summary_3:              "[миссия %d] Средняя скорость: %v квадрат/сек",
	UnsupportedACMEVersion:      "Эта версия структуры ACME не поддерживается (поддерживается только версия файла acme 1.2)",
	Warning_ACME_Deprecated:     "Предупреждение - `acme' является устаревшим и был удален, пожалуйста, перейдите на формат BDX вместо него.",
	Warning_Schem_Deprecated:    "Предупреждение - `schem' является устаревшим и был удален, пожалуйста, перейдите на формат BDX вместо него.",
	Warning_UserHomeDir:         "Предупреждение - Домашний каталог текущего пользователя не может быть получен, он будет установленhomedir=\".\";\n",
}
