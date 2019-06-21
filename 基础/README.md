Go语言学习之路
~~~~数据结构总结
数组：  像php中不设定key的数组,定长
	声明 var array [10] int
		 array := [2] int {1,2}
 结构体： 像php中的对象还能有方法
 	声明 type Books struct {
			title string
			author string
			subject string
			book_id int
            Tes 	map[int]int   //用new,{},var 要选make
		}
		book := Books{title:"解忧杂货",author:"小高"}
 切片(动态数组)，不定长
    声明 var slice [] int 
         s := make([] int, 3, 4)
         s := [] int {1,2,3} 
         
         array := [30] int {1,2,3}
         s := array[:]
 集合 ：像php中的设定key的数组
    声明 var country map[string]string
        country  = make(map[string]string)
        
        countrymap := map[string]string{}
         
         
         