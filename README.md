
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
<p>Маршрут: http://[domainname:port]/user </p>

<table class="table1" cellpadding="8">
            <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
            <tr>
                <td>1</td> <td>name</td> <td>Имя. </td> 
            </tr>
			 <tr>
                <td>2</td> <td>surname</td> <td>Фамилия. </td>
            </tr>
			   <tr>
                <td>3</td> <td>midlename</td> <td>Отчество. </td>
            </tr>
			   <tr>
                <td>4</td> <td>gender</td> <td>Пол.</td>
            </tr>
			   <tr>
                <td>5</td> <td>age</td> <td>Возраст </td> 
				</tr>
		    </table>
		
		
<h3>Ответ 200OK</h3>
<table class="table1" cellpadding="8">
            <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			<tr>
			    <td>1</td> <td>uuid</td> <td>uuid созданного пользователя </td> 
            </tr>
            <tr>
                <td>2</td> <td>name</td> <td>Имя. </td> 
            </tr>
			 <tr>
                <td>3</td> <td>surname</td> <td>Фамилия. </td>
            </tr>
			   <tr>
                <td>4</td> <td>midlename</td> <td>Отчество. </td>
            </tr>
			    <tr>
                <td>5</td> <td>fio</td> <td>surname+name+middlename</td> 
            </tr>
			   <tr>
                <td>6</td> <td>gender</td> <td>Пол.</td>
            </tr>
			   <tr>
                <td>7</td> <td>age</td> <td>Возраст </td> 
		    </tr>
        </table>
<h3>Другие варианты ответов 500, 400</h3>

<h3 align="left"> Изменить пользователя</h3>
<p>Метод: PUT</p>
<p>Маршрут: http://[domainname:port]/users/{uuid пользователя} </p>
<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
		    <tr>
                <td>1</td> <td>name</td> <td>Имя. </td> 
            </tr>
			 <tr>
                <td>2</td> <td>surname</td> <td>Фамилия. </td>
            </tr>
			   <tr>
                <td>3</td> <td>midlename</td> <td>Отчество. </td>
            </tr>
			    <tr>
                <td>4</td> <td>fio</td> <td>surname+name+middlename</td> 
            </tr>
			   <tr>
                <td>5</td> <td>gender</td> <td>Пол.</td>
            </tr>
			   <tr>
                <td>6</td> <td>age</td> <td>Возраст </td>  
			  </tr>
        </table>
<h3>204, 500, 400</h3>

<h3 align="left">Удалить пользователя</h3>
<p>Метод: DELETE</p>
<p>Маршрут: http://[domainname:port]/users/{uuid пользователя} </p>
<h3>Варианты ответов 204, 500, 400</h3> 



 <h3 align="left">Создать продукт</h3>

<p>Метод: POST</p>
<p>Маршрут: http://[domainname:port]/product</p>

<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			 <tr>
                <td>1</td> <td>description</td> <td>Описание товара. </td> 
            </tr>
			<tr>
                <td>2</td> <td>price</td> <td>Цена товара. </td> 
            </tr>
			<tr>
                <td>3</td> <td>currency</td> <td>Валюта </td> 
            </tr>
			<tr>
                <td>4</td> <td>left_in_stock</td> <td>Остаток на складе. </td> 
            </tr>
</table>			
<h3>Ответ 200OK</h3>

<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			 <tr>
                <td>1</td> <td>uuid</td> <td>uuid созданного товара. </td> 
            </tr>
			 <tr>
                <td>2</td> <td>description</td> <td>Описание товара. </td> 
            </tr>
			<tr>
                <td>3</td> <td>price</td> <td>Цена товара. </td> 
            </tr>
			<tr>
                <td>4</td> <td>currency</td> <td>Валюта </td> 
            </tr>
			<tr>
                <td>5</td> <td>left_in_stock</td> <td>Остаток на складе. </td> 
            </tr>
</table>
<h3>Другие варианты ответов 500, 400</h3>
<h3 align="left"> Изменить продукт</h3>
<p>Метод: PUT</p>
<p>Маршрут: http://[domainname:port]/products/{uuid продукта} </p>
<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			 <tr>
                <td>1</td> <td>description</td> <td>Описание товара. </td> 
            </tr>
			<tr>
                <td>2</td> <td>price</td> <td>Цена товара. </td> 
            </tr>
			<tr>
                <td>3</td> <td>currency</td> <td>Валюта </td> 
            </tr>
			<tr>
                <td>4</td> <td>left_in_stock</td> <td>Остаток на складе. </td> 
            </tr>
</table>
<h3>Другие варианты ответов 204, 500, 400</h3>
<h3 align="left">Удалить продукт</h3>
<p>Метод: DELETE</p>
<p>Маршрут: http://[domainname:port]/products/{uuid продукта} </p>
<h3>Варианты ответов 204, 500, 400</h3> 


 <h3 align="left">Создать заказ</h3>

<p>Метод: POST</p>
<p>Маршрут: http://[domainname:port]/order/{uuid пользователя}</p>

<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			 <tr>
                <td>1</td> <td>product</td> <td>JSON массив из uuid продукта</td> 
            </tr>
		
</table>			
<h3>Ответ 200OK</h3>

<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			 <tr>
                 <td>1</td> <td>product</td> <td>JSON массив из uuid продукта</td> 
            </tr>
			 <tr>
                 <td>2</td> <td>uuid</td> <td>JSON массив из uuid ордеров</td>  
            </tr>
			
</table>
<p>Заказ представялет собой совокупность ордеров. Может иметь несколько продуктов. Ордер представляют собой номер заказа к которому он относиться и может
иметь только один продукт. Ордер нельзя изменить только создать и удалить. При возврате ответа важен порядок product и uuid , так как i-й элемент из productt 
соответсует i-му элемиенту uuid</p> 
 
<h3>Другие варианты ответов 500, 400</h3>

<h3 align="left"> Изменить заказ</h3>
<p>Метод: PUT</p>
<p>Маршрут: http://[domainname:port]/order/{uuid пользователя} </p>
<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			 <tr>
                 <td>1</td> <td>product</td> <td>JSON массив из uuid продукта</td> 
            </tr>
			 <tr>
                 <td>2</td> <td>uuid</td> <td>JSON массив из uuid ордеров</td>  
            </tr>
			<tr>
                 <td>3</td> <td>id</td> <td>Номер заказа пользователя</td>  
            </tr>		
</table>
<h3>Ответ 200OK</h3>

<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
			 <tr>
                 <td>1</td> <td>product</td> <td>JSON массив из uuid продукта</td> 
            </tr>
			 <tr>
                 <td>2</td> <td>uuid</td> <td>JSON массив из uuid ордеров</td>  
            </tr>
			<tr>
                 <td>3</td> <td>id</td> <td>Номер заказа пользователя</td>  
            </tr>	
			
</table>
<h3>500, 400</h3>

<h3 align="left">Удалить продукт</h3>
<p>Метод: DELETE</p>
<p>Маршрут: http://[domainname:port]/order/{uuid пользователя} </p>

<table class="table1" cellpadding="8">
           <tr class="table_font1">
                <td >№</td> <td>Наименование параметра</td> <td>Значение</td>
            </tr>
				<tr>
                 <td>1</td> <td>id</td> <td>Номер заказа пользователя</td>  
            </tr>	
			
</table>
<h3>Варианты ответов 204, 500, 400</h3> 


