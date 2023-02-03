
tags: #Go 

---

Всем привет, в этой статье мы рассмотрим небольшой пример кода, написанный для сбора данных с сайта на популярном сегодня языке программирования Go.

Эта статья будет особенно интересна тем, кто много слышал о Go, но пока еще не попробовал его самостоятельно. 

Для одного из внутренних проектов, нам нужно было собрать данные о фильмах, которые идут в прокате и расписании показов, сейчас рассмотрим первую (простейшую) версию парсера, с которой все и началось.
<a name="habracut"></a>
Для тех, кому лень читать код статьи, сразу привожу ссылку на <a href="https://github.com/noorsoft-mobile/go-cinema-parser" rel="nofollow">репозиторий</a>.

В первой своей версии парсер умел собирать лишь следующую информацию:


- О сеансах в одном кинотеатре,

- Детальное описание,

- Название,

- Киностудия,



И сохранять ее в JSON. Первоначально необходимо выбрать подходящую библиотеку для парсинга.

Google выдает большое количество вариантов по запросу `golang web scraping`, многие из них отражены в этом <a href="https://github.com/lorien/awesome-web-scraping/blob/master/golang.md" rel="nofollow">списке</a>, советую ознакомиться с ним, свой же выбор я остановил на <a href="https://github.com/geziyor/geziyor" rel="nofollow">geziyor</a>, так как он поддерживает JS Rendering (который, кстати мы не будем использовать в этом примере, однако эта фича бывает очень полезной при парсинге сайтов) и довольно прост в использовании.

Итак, библиотека выбрана, следующий шаг — установить ее и начать использовать ее в коде.
Установка библиотеки крайне проста:


```go
go get -u github.com/geziyor/geziyor
```

Теперь перейдем к написанию кода.

Внутри тела главной функции вызовем парсер, передадим ему URL страницы, с которой начнем сбор данных и укажем, что хотим экспортировать результат в JSON файл:


```go

func main() { 
    geziyor.NewGeziyor(&amp;geziyor.Options{ 
        StartURLs: []string{&#34;https://kinoteatr.ru/raspisanie-kinoteatrov/city/#&#34;}, 
        ParseFunc: parseMovies, 
        Exporters: []export.Exporter{&amp;export.JSON{}}, 
    }).Start() 
}

```

Начало положено, но не хватает логики сбора данных, для этого нужно реализовать функцию <b>parseMovies</b>.

Логика сбора будет следующей:


- Поиск блока, содержащего информацию о фильме,

- Сбор информации о всех сеансах внутри этого блока,

- Сбор названия фильма и киностудии,

- Сбор ссылки на страницу с подробной информацией о фильме,

- Сбор описания с этой страницы




#### Перейдем к реализации этой функции


Здесь выбираются все блоки, содержащие информацию о фильме, для дальнейшей их обработки.


```go
func parseMovies(g *geziyor.Geziyor, r *client.Response) {
    r.HTMLDoc.Find(&#34;div.shedule_movie&#34;).Each(func(i int, s *goquery.Selection) { 
```

Таким образом собирается информация о сеансах, и тут же приводится в удобный человеку вид (убираем лишние пробелы и отступы на новую строку).


```go

var sessions = strings.Split(s.Find(&#34;.shedule_session_time&#34;).Text(), &#34; \n &#34;) 
sessions = sessions[:len(sessions)-1] 

for i := 0; i &lt; len(sessions); i++ { 
    sessions[i] = strings.Trim(sessions[i], &#34;\n &#34;) 
}

```

Этот блок кода отвечает за получение страницы с детальной информацией о фильме и получение его описания.


```go

if href, ok := s.Find(&#34;a.gtm-ec-list-item-movie&#34;).Attr(&#34;href&#34;); ok {
    g.Get(r.JoinURL(href), func(_g *geziyor.Geziyor, _r *client.Response) {
        description = _r.HTMLDoc.Find(&#34;span.announce p.movie_card_description_inform&#34;).Text() 
        description = strings.ReplaceAll(description, &#34;\t&#34;, &#34;&#34;) 
        description = strings.ReplaceAll(description, &#34;\n&#34;, &#34;&#34;) 
        description = strings.TrimSpace(description) 

```

Так вызывается API для экспортирования результатов в JSON файл.


```go

g.Exports &lt;- map[string]interface{}{ 
    &#34;title&#34;:        strings.TrimSpace(s.Find(&#34;span.movie_card_header.title&#34;).Text()), 
    &#34;subtitle&#34;:    strings.TrimSpace(s.Find(&#34;span.sub_title.shedule_movie_text&#34;).Text()), 
    &#34;sessions&#34;:    sessions, 
    &#34;description&#34;: description, 
}

```

Ура, все готово! Осталось только объединить написанные блоки кода воедино и запустить парсинг.

Так выглядит процесс работы парсера, видим в терминале сообщения об успешном получении страниц, это удобно.

<img src="https://habrastorage.org/r/w1560/webt/ab/ee/zo/abeezot2smu6-n26c31884sg5tg.png" data-src="https://habrastorage.org/webt/ab/ee/zo/abeezot2smu6-n26c31884sg5tg.png"/>

А так выглядит результат парсинга.

<img src="https://habrastorage.org/r/w1560/webt/0f/qy/3v/0fqy3vtnnv727n5xbjy-jju9zh4.png" data-src="https://habrastorage.org/webt/0f/qy/3v/0fqy3vtnnv727n5xbjy-jju9zh4.png"/>

Спасибо за чтение статьи, любите программирование на Go.