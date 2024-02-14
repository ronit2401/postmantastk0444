package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type MenuItem struct {
	ItemNo int    `json:"item_no"`
	Item   string `json:"item"`
}

type DayMenu struct {
	Day   string        `json:"day"`
	Meals [3][]MenuItem `json:"meals"`
}

type matrix struct {
	itemno [7][3]int
	items  [7][3]string
}

func q1(itemno [7][3]int, items [7][3]string) {
	fmt.Println("enter the day->(1/2/3/4/5/6/7) and meal(b/l/d)->(1/2/3)")
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	a = a - 1
	b = b - 1
	fmt.Println("the items are :", items[a][b])
}

func q2(itemno [7][3]int, items [7][3]string) {
	fmt.Println("enter the day->(1/2/3/4/5/6/7) and meal(b/l/d)->(1/2/3)")
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	a = a - 1
	b = b - 1
	fmt.Println("the number of items in the menu are :", itemno[a][b])
}

func q3(itemno [7][3]int, items [7][3]string) {
	fmt.Println("enter the day->(1/2/3/4/5/6/7) and meal(b/l/d)->(1/2/3) and item to check")
	var a int
	var b int
	var c string
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	a = a - 1
	b = b - 1
	if strings.TrimSpace(strings.ToLower(items[a][b])) == strings.TrimSpace(strings.ToLower(c)) {
		fmt.Println("item is in the menu")
	} else {
		fmt.Println("item not in the menu")
	}

}
func ConvertAndSaveMenuToJSON(itemno [7][3]int, items [7][3]string) error {
	menu := make([]DayMenu, len(items))
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(items); i++ {
		menu[i].Day = days[i]
		for j := 0; j < len(items[i]); j++ {
			menu[i].Meals[j] = parseMenuItems(items[i][j], itemno[i][j])
		}
	}

	// Convert menu to JSON
	menuJSON, err := json.MarshalIndent(menu, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Write JSON to file
	file, err := os.Create("menu.json")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(menuJSON)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %v", err)
	}

	fmt.Println("Menu data saved to menu.json")
	return nil
}

func parseMenuItems(menuString string, itemCount int) []MenuItem {
	items := strings.Split(menuString, " ")
	menuItems := make([]MenuItem, 0, len(items))
	for _, item := range items {
		menuItems = append(menuItems, MenuItem{
			ItemNo: itemCount,
			Item:   item,
		})
	}
	return menuItems
}

func main() {

	itemno := [7][3]int{
		{7, 8, 6},
		{8, 7, 4},
		{8, 4, 4},
		{9, 8, 5},
		{4, 8, 6},
		{7, 7, 2},
		{8, 5, 7},
	}
	items := [7][3]string{
		{
			"choice of egg cornflakes bread +jam poha grapes tea+ coffee milk",
			"disco papad kadhi pakoda madrasi aloo veg khichdi chapati plain  rice",
			"mirch ke tapore/ lahsun chutney aloo rassa dal fry bati/ chapati plain  rice rasam",
		},
		{
			"no egg cornflakes bread +jam aloo paratha curd pickel kinu tea+ coffee milk",
			"bhel poori sweet potato dry veg kadhai dal maharani plain  rice sambar sweet lassi",
			"lemon + onion khatta meetha pumpkin chole  masala hot choco milk",
		},
		{
			"choice of egg cornflakes bread +jam bread pakoda tomato ketchup sweet daliya tea+ coffee milk",
			"chapati veg polao pineapple raita choco burfi",
			"aloo chana chaat tawa veg (no karela) dal palak chapati",
		},
		{
			"choice of egg cornflakes bread +jam mutter kulcha kulcha cut onion lemon papaya tea+ coffee milk",
			"green salad gajar mutter gatta curry channa dal tadka chapati plain  rice sambar rasna",
			"green salad veg korma butter dal tadka chapati gulab jamun",
		},
		{
			"choice of egg cornflakes tea+ coffee milk",
			"onion salad baingan bhartaha rajma masala dal fry chapati plain  rice bathua raita shikanji",
			"mix salad veg manchurian black masoor dal egg fried rice / veg fried rice chapati baliushahi",
		},
		{
			"choice of egg cornflakes bread +jam suji upma adrak chutney tea+ coffee milk",
			"green chutney french fries ragda mutter curd /butter mix veg paratha plain  rice sambar",
			"green salad sarsoon ka saag",
		},
		{
			"choice of egg cornflakes bread +jam masala dosa sambhar grapes tea+ coffee milk",
			"onion salad paneer lababdar dhaba chicken naan / chapati biryani rice",
			"pasta salad methi malai mutter dal lasooni plain rice chapati rasam hot & sour soup",
		},
	}

	j := 1
	for j != 0 {
		fmt.Println("enter 1 for task 1, 2 for task 2, 3 for task 3, 4 for task 4, and 5 to quit")
		var b int
		fmt.Scan(&b)
		switch b {
		case 1:
			q1(itemno, items)
		case 2:
			q2(itemno, items)
		case 3:
			q3(itemno, items)
		case 4:
			{
				err := ConvertAndSaveMenuToJSON(itemno, items)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		case 5:
			j = 0
		default:
			fmt.Println("wrong input")
		}
	}

}
