API Создания реферальных кодов, хранения и регистрация по ним.

Функциональные требования:
  - Регистрация и аутентификация пользователя (JWT, Oauth 	2.0);
  - Аутентифицированный 	пользователь должен иметь возможность 	создать или удалить свой реферальный код.Одновременно может быть активен только 1 код. При создании кода обязательно 	должен быть задан его срок годности;
  - Возможность получения реферального кода по email адресу реферера;
  - Возможность регистрации по реферальному коду в 	качестве реферала;
  - Получение 	информации о рефералах по id реферера;
  - UI документация (Swagger/ReDoc)


Запуск под средствам docker-compose
  <pre><code>docker-compose up -d</code></pre>

Все пременные окружения установленны в compose файле

Настроена сеть между контейнерами 192.168.200.0/24
ip контейнера с приложением 192.168.200.104

Описание API доступно в спецификации
<pre><code>ReferralsAPI.postman_collection.json</code></pre>
