package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура, описывающая список расходов и имеющая идентификатор
type ExpenseList struct {
	NextID   int 		`json:next_id`
	Expenses []Expense 	`json:expenses`
}
// Структура расхода
type Expense struct {
	ID     int 			`json: id`
	Name   string		`json: name`
	Desc   string		`json: descrition`
	Amount float64		`json: amount`
}

var expenseList ExpenseList

// Метод для добавления расхода
func (el *ExpenseList) AddExpense(scanner *bufio.Scanner) error {
	fmt.Print("Введите название расхода: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		return fmt.Errorf("название не может быть пустым")
	}

	fmt.Print("Введите описание расхода: ")
	scanner.Scan()
	desc := strings.TrimSpace(scanner.Text())
	if desc == "" {
		return fmt.Errorf("описание не может быть пустым")
	}

	fmt.Print("Введите потраченную сумму: ")
	scanner.Scan()
	amount := strings.TrimSpace(scanner.Text())
	if amount == "" {
		return fmt.Errorf("сумма не может быть пустой")
	}

	amountF, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("Ошибка конвертации: %v", err)
		
	}

	newExpense := Expense{
		ID:     el.NextID ,
		Name:   name,
		Desc:   desc,
		Amount: amountF,
	}
	el.Expenses = append(el.Expenses, newExpense)
	el.NextID++
	fmt.Printf("Расход '%s' добавлен с номером %d.\n", name, newExpense.ID)
	return nil
}

// Метод для удаления расхода
func (el *ExpenseList) deleteExpense(scanner *bufio.Scanner) error{
	fmt.Println("Введите ID расхода, который вы собираетесь удалить:")
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text())
	if idStr == ""{
		return fmt.Errorf("Вы ничего не ввели!")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0{
		return fmt.Errorf("Ошибка: %v", err)
	}

	for i, expense := range el.Expenses{
		if expense.ID == id{
			fmt.Printf("Вы уверены, что хотите его удалить '%s'? (да/нет):", expense.Name)
			scanner.Scan()
			confirm := strings.TrimSpace(scanner.Text())
			if confirm == "да" {
				el.Expenses = append(el.Expenses[:i], el.Expenses[i+1:]...) // Данная конструкция используется для удаления элемента из слайса
				fmt.Printf("Расход с ID %d удален.", id)
				return err
			} else{
				fmt.Println("Удаление отменено")
				return err
			} 
		}
	}
	return nil
}

// Метод для обновления информации о расходе
func (el *ExpenseList) updateExpense(scanner *bufio.Scanner) error{
	fmt.Println("Введите ID расхода, которые вы собираетесь обновить: ")
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text())
	if idStr == ""{
		return fmt.Errorf("Вы ничего не ввели!")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0{
		return fmt.Errorf("Ошибка: ", err)
	}

	for i, expense := range el.Expenses{
		if expense.ID == id{
			fmt.Println("Что конкретно вы хотите обновить?")
			scanner.Scan()
			choice := strings.TrimSpace(scanner.Text())
			switch choice{
				case "Название": 
					fmt.Println("Введите новое название: ")
					scanner.Scan()
					newName := strings.TrimSpace(scanner.Text())
					el.Expenses[i].Name = newName
					fmt.Println("Название успешно обновлено!")
				case "Описание": 
					fmt.Println("Введите новое описание: ")
					scanner.Scan()
					newDesc := strings.TrimSpace(scanner.Text())
					el.Expenses[i].Desc = newDesc
					fmt.Println("Описание успешно обновлено!")
				case "Сумма": 
					fmt.Println("Введите новую сумму: ")
					scanner.Scan()
					newAmountStr := strings.TrimSpace(scanner.Text())
					newAmount, err := strconv.Atoi(newAmountStr)
					if err != nil{
						fmt.Errorf("Ошибка: ", err)
					}
					el.Expenses[i].Amount = float64(newAmount)
					fmt.Println("Сумма успешно обновлена!")
				default:
					fmt.Println("Неизвестная команда!")
				}
		}
	}
	return nil
}

// Функция для поиска конкретного расхода
func searchExpense(scanner *bufio.Scanner) error{
	fmt.Println("Введите ID расхода, который вы хотите найти: ")
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text())
	if idStr == ""{
		return fmt.Errorf("Вы ничего не ввели!")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0{
		return fmt.Errorf("Ошибка: ", err)
	}
	fmt.Println("\nРезультат поиска:")
	count := 0
	for _, expense := range expenseList.Expenses{
		if expense.ID == id{
			fmt.Printf("ID: %d | Название: %s | Описание: %s | Сумма: %g\n", 
			expense.ID, expense.Name, expense.Desc, expense.Amount)
		count++
		}
	}
	if count == 0 {
		fmt.Println("Расходы не найдены.")
	}
	return nil
}

// Функция для подсчета суммы всех расходов
func totalSummary(){
	count := 0.0
	for _, expense := range expenseList.Expenses{
		count += float64(expense.Amount)
	}
	fmt.Println(count)
}
// Функция для вывода информации о расходе
func printExpenses() {
	fmt.Println("\nРасходы:")
	count := 0
	for _, expense := range expenseList.Expenses {
		fmt.Printf("ID: %d | Название: %s | Описание: %s | Сумма: %g\n", 
			expense.ID, expense.Name, expense.Desc, expense.Amount)
		count++
	}
	if count == 0 {
		fmt.Println("Расходы не найдены.")
	}
}



// Функция менюшки
func menu() {
	fmt.Println("CLI Трекер расходов")
	fmt.Println("1. Добавить")
	fmt.Println("2. Обновить")
	fmt.Println("3. Просмотреть")
	fmt.Println("4. Удалить")
	fmt.Println("5. Найти")
	fmt.Println("6. Просмотреть сумму за все время")
	fmt.Println("7. Сохранить расходы в json файл")
	fmt.Println("8. Выйти")
	fmt.Print("Выберите действие: ")
}

// Функция для сохранение расходов в JSON
func saveToFile() {
	file, err := os.Create("expenses.json")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(expenseList, "", "  ")
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}
	fmt.Println("Задачи сохранены в expenses.json")
}


// Входная точка программы
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		menu()
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		
		switch choice {
		case "1":
			if err := expenseList.AddExpense(scanner); err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "2":
			if err := expenseList.updateExpense(scanner); err != nil{
				fmt.Println("Ошибка:", err)
			}
		case "3":
			printExpenses()
		case "4":
			if err := expenseList.deleteExpense(scanner); err != nil{
				fmt.Println("Ошибка: ", err)
			}
		case "5":
			searchExpense(scanner)
		case "6":
			totalSummary()
		case "7": 
			saveToFile()
		case "8":
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}