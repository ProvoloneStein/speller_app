<h1>Работа с приложение</h1>

/api/v1/text для работы с одним текстом
принимает json формата {"text": "Ваш текст"}

/api/v1/texts для работы с несколькими текстами (для работы с данными из ТЗ)
принимает json формата {"text": ["Ваш текс", "Еще один ващ текст"...]}

возвращает массив из позиции слова, строки в тексе, слова и вариантов его исправления

по стандарту запускатется в dev, для запуска в прод необходимо добавить в окружения переменные из prod.env
<h1>#TODO</h1> 
Имеет смысл менеджить ошибки Яндекс - спеллера 
