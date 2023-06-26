package graphmodels

import (
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	uri = "bolt://localhost:7687"
)

type Movie struct {
	Title string `json:"title"`
}

type Person struct {
	Name string `json:"name"`
}

func CreateDriver(uri, username, password string) (neo4j.Driver, error) {
	return neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
}

func CreateSession(driver neo4j.Driver, accessMode neo4j.AccessMode) (neo4j.Session, error) {
	return driver.NewSession(neo4j.SessionConfig{AccessMode: accessMode}), nil
}

func NodeQuery(session neo4j.Session, query string, params map[string]interface{}) (neo4j.Result, error) {
	return session.Run(query, params)
}

// String类型 转成 map[string]interface{}类型
func StringToMap(content string) map[string]interface{} {
	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &resMap)
	if err != nil {
		fmt.Println("string转map失败", err)
	}
	return resMap
}

func Graph() (ret []Movie) {
	d, err := CreateDriver(uri, "", "")
	if err != nil {
		panic(err)
	}
	s, err2 := CreateSession(d, neo4j.AccessModeRead)
	if err2 != nil {
		panic(err2)
	}
	//query := "match (n:Person) return n"
	query := "MATCH (a:Movie)-[:actor]->(b:Person {name:'三浦贵大'}) RETURN a"
	r, err3 := NodeQuery(s, query, nil)
	//fmt.Println(r, err3) &{0xc0000a6000} <nil>

	if err3 != nil {
		panic(err3)
	}
	var movies []Movie
	for r.Next() {
		record := r.Record()
		values := record.Values
		//fmt.Println(values) [{1 1 [Person] map[name:三浦贵大]}]

		for _, value := range values {
			movie := value.(neo4j.Node).GetProperties()
			fmt.Println(movie)
			movies = append(movies, Movie{Title: movie["title"].(string)})
			println(movie["title"].(string))
		}
	}
	return movies
}

// 根据名字查找人
func Find_Node(name string) (ret []Person) {
	//建立连接
	d, err := CreateDriver(uri, "", "")
	if err != nil {
		panic(err)
	}
	//建立操作对象
	s, err2 := CreateSession(d, neo4j.AccessModeRead)
	if err2 != nil {
		panic(err2)
	}

	//name_map := StringToMap(name)
	r, err3 := NodeQuery(s, "Match (a:Person{name:$name}) return a", map[string]any{"name": name})

	if err3 != nil {
		panic(err3)
	}
	var persons []Person
	for r.Next() {
		record := r.Record()
		values := record.Values
		fmt.Println(values) // [{1 1 [Person] map[name:三浦贵大]}]

		for _, value := range values {
			person := value.(neo4j.Node).GetProperties()
			fmt.Println(person)
			persons = append(persons, Person{Name: person["name"].(string)})
			println(person["name"].(string))
		}
	}
	return persons
}

// 根据title和id增人
func InsertPernson(title string, name string) (ret []Person) {
	//建立连接
	d, err := CreateDriver(uri, "", "")
	if err != nil {
		panic(err)
	}
	//建立操作对象
	s, err2 := CreateSession(d, neo4j.AccessModeRead)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("111", title, name)
	//name_map := StringToMap(name)
	r, err3 := NodeQuery(s, "CREATE (n:Person {title:$title, name:$name}) "+
		"RETURN n ", map[string]any{"title": title, "name": name})

	if err3 != nil {
		panic(err3)
	}
	var persons []Person
	for r.Next() {
		record := r.Record()
		values := record.Values
		fmt.Println(values) // [{1 1 [Person] map[name:三浦贵大]}]

		for _, value := range values {
			person := value.(neo4j.Node).GetProperties()
			fmt.Println(person)
			persons = append(persons, Person{Name: person["name"].(string)})
			println(person["name"].(string))
		}
	}
	return persons
}

// 根据title和id删人
func DeletePernson(title string, id int) (ret []Person) {
	//建立连接
	d, err := CreateDriver(uri, "", "")
	if err != nil {
		panic(err)
	}
	//建立操作对象
	s, err2 := CreateSession(d, neo4j.AccessModeRead)
	if err2 != nil {
		panic(err2)
	}

	//name_map := StringToMap(name)
	//r, err3 := NodeQuery(s, "Match (a:Person{title:$title,id:$id}) Delete a", map[string]any{"title": title, "id": id})
	r, err3 := NodeQuery(s, "Match (a:Person{title:$title}) "+
		"Where id(a)=$id Delete a", map[string]any{"title": title, "id": id})
	if err3 != nil {
		panic(err3)
	}
	var persons []Person
	for r.Next() {
		record := r.Record()
		values := record.Values
		fmt.Println(values) // [{1 1 [Person] map[name:三浦贵大]}]

		for _, value := range values {
			person := value.(neo4j.Node).GetProperties()
			fmt.Println(person)
			persons = append(persons, Person{Name: person["name"].(string)})
			println(person["name"].(string))
		}
	}
	return persons
}

// 根据title查找人
func SearchPernson(title string, id int) (ret []Person) {
	//建立连接
	d, err := CreateDriver(uri, "", "")
	if err != nil {
		panic(err)
	}
	//建立操作对象
	s, err2 := CreateSession(d, neo4j.AccessModeRead)
	if err2 != nil {
		panic(err2)
	}

	//name_map := StringToMap(name)
	r, err3 := NodeQuery(s, "Match (a:Person{title:$title}) "+
		"Where id(a)=$id return a", map[string]any{"title": title, "id": id})

	if err3 != nil {
		panic(err3)
	}
	var persons []Person
	for r.Next() {
		record := r.Record()
		values := record.Values
		fmt.Println(values) // [{1 1 [Person] map[name:三浦贵大]}]

		for _, value := range values {
			person := value.(neo4j.Node).GetProperties()
			fmt.Println(person)
			persons = append(persons, Person{Name: person["name"].(string)})
			println(person["name"].(string))
		}
	}
	return persons
}

// 根据title和id和attribute改人
func ModifyPernson(title string, id int, attributes map[string]string) (ret []Person) {
	//建立连接
	d, err := CreateDriver(uri, "", "")
	if err != nil {
		panic(err)
	}
	//建立操作对象
	s, err2 := CreateSession(d, neo4j.AccessModeRead)
	if err2 != nil {
		panic(err2)
	}

	r, err3 := NodeQuery(s, "Match (n:Person{title:$title}) "+
		"Where id(n)=$id Set n.age=$age,n.sex=$sex,n.city=$city return n",
		map[string]any{"title": title, "id": id, "age": attributes["age"], "sex": attributes["sex"]})

	if err3 != nil {
		panic(err3)
	}
	var persons []Person
	for r.Next() {
		record := r.Record()
		values := record.Values
		fmt.Println(values) // [{1 1 [Person] map[name:三浦贵大]}]

		for _, value := range values {
			person := value.(neo4j.Node).GetProperties()
			fmt.Println(person)
			persons = append(persons, Person{Name: person["name"].(string)})
			println(person["name"].(string))
		}
	}
	return persons
}
