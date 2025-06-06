package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB

// Initialize the database connection
func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	//　AutoMigrate the Todo model
	db.AutoMigrate(&Todo{})
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func getTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func deleteTodo(c *gin.Context) {
	// パラメータからIDを取得
	idPram := c.Param("id")
	// 文字を整数に変換
	id, err := strconv.Atoi(idPram)
	// エラー処理
	// 変換に失敗した場合はエラーレスポンスを返す
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// IDを使ってデータベースからTodoを取得
	var todo Todo
	// データベースからTodoを検索
	// Firstメソッドを使って、IDに一致する最初のレコードを取得
	// エラーが発生した場合は、404エラーを返す
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	// Todoが見つかった場合は、削除処理を行う
	db.Delete(&todo)
	// 削除が成功した場合は、成功レスポンスを返す
	// HTTPステータスコード200とメッセージを返す
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func updateTodo(c *gin.Context) {
	isParam := c.Param("id")
	id, err := strconv.Atoi(isParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 値を更新
	todo.Title = input.Title
	todo.Completed = input.Completed

	db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func main() {
	r := gin.Default()

	// CORS設定を追加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	initDB()

	r.POST("/todos", createTodo)
	r.GET("/todos", getTodos)
	r.DELETE("/todos/:id", deleteTodo)
	r.PUT("/todos/:id", updateTodo)
	r.Run(":8080") // Run on port 8080
}
