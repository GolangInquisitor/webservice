<h2>Test web service</h2>
<p></p>
<h2>Instruction</h2>
<li> Install PostgreSQL</li>
<li> psql  -U postgres -d <base> /scripts/postgres/scoltest.sql</li>
<li> build project cmd\server_app.main.go</li>
<li>place the file config.yml next to the executable file </li>
<li> change the settings in config.yml if necessary</li>
 <li> run
 <h2>API</h2>
 <h3 align="left">Создать пользователя</h3>

<p>Метод: POST</p>
<p>Маршрут: http://[domainname:port]/create </p>
<table class="table1" cellpadding="8">
            <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
            <tr>
                <td>1</td> <td>token</td> <td> Строка JWT токена полученного при аутентификации </td>
            </tr>
            <tr>
                <td>2</td> <td>cmd_arg</td><td> user </td>
            </tr>
            <tr>
                <td>3</td> <td>name</td> <td>Имя. </td> 
            </tr>
			 <tr>
                <td>4</td> <td>lname</td> <td>Фамилия. </td>
            </tr>
			   <tr>
                <td>5</td> <td>patr</td> <td>Отчество. </td>
            </tr>
			   <tr>
                <td>6</td> <td>post</td> <td>Должность в организации.</td>
            </tr>
			   <tr>
                <td>7</td> <td>log</td> <td>Логин </td>
            </tr>
			   <tr>
                <td>8</td> <td>pass</td><td>Пароль</td>
            </tr>
			   
        </table>
<p>Новый пользователь всегда создается не активным. Чтобы пользователь мог входить в сеть, надо его активировать после создания.</p>
<h3 align="left">Изменить пользователя</h3>
<p>Метод: POST</p>
<p>Маршрут: http://[domainname:port]/change </p>
<table class="table1" cellpadding="8">
            <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
            <tr>
                <td>1</td> <td>token</td> <td> Строка JWT токена полученного при аутентификации </td>
            </tr>
            <tr>
                <td>2</td> <td>cmd_arg</td><td> user </td>
            </tr>
			<tr>
                <td>3</td> <td>id</td> <td>[id пользователя].  Должен быть известен, менять нельзя</td> 
            </tr>
			 <tr>
                <td>3</td> <td>name</td> <td>Имя. Если менять не надо отравляем старое значение.</td> 
            </tr>
			 <tr>
                <td>4</td> <td>lname</td> <td>Фамилия. Если менять не надо отравляем старое значение.</td>
            </tr>
			   <tr>
                <td>5</td> <td>patr</td> <td>Отчество. Если менять не надо отравляем старое значение.</td>
            </tr>
			   <tr>
                <td>6</td> <td>post</td> <td>Должность в организации.Если менять не надо отравляем старое значение.</td>
            </tr>
			   <tr>
                <td>7</td> <td>log</td> <td>Логин. Если менять не надо отравляем старое значение.</td>
            </tr>
			   <tr>
                <td>8</td> <td>pass</td><td>Пароль. Если менять не надо отравляем старое значение.</td>
            </tr>
			<tr>
                <td>3</td> <td>active</td> <td>либо "true" либо "false". Если менять не надо отравляем старое значение. </td> 
            </tr>
			   
        </table>


<h3 align="left">Удалить пользователя</h3>
<p>Метод: POST</p>
<p>Маршрут: http://[domainname:port]/delete </p>
<table class="table1" cellpadding="8">
            <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
            <tr>
                <td>1</td> <td>token</td> <td> Строка JWT токена полученного при аутентификации </td>
            </tr>
            <tr>
                <td>2</td> <td>cmd_arg</td><td> user </td>
            </tr>
			<tr>
                <td>3</td> <td>id</td> <td>[id пользователя].  Должен быть известен</td> 
            </tr>		
			   
        </table>
 