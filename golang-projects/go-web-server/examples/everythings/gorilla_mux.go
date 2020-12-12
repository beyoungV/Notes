URL:
/books/go-programming-blueprint/page/10
//该URL具有两个动态段：
books（go-programming-blueprint）
page（10）

//使请求处理程序与上述URL匹配，您可以使用URL模式中的占位符替换动态分段，如下所示：
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    // get the book
    // navigate to the page
})

//从这些段中获取数据。该包带有功能mux.Vars(r)它接受http.Request作为参数并返回地图上的节段。
func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    vars["title"] // the book title slug
    vars["page"] // the page
}
