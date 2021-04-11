package main

import "fmt"

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)package main

	import (
		"fmt"
		"log"
	)

	func main() {
		vendingMachine := newVendingMachine(1, 10)
		err := vendingMachine.requestItem()
		if err != nil {
			log.Fatalf(err.Error())
		}
		err = vendingMachine.insertMoney(10)
		if err != nil {
			log.Fatalf(err.Error())
		}
		err = vendingMachine.dispenseItem()
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println()
		err = vendingMachine.addItem(2)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println()

		err = vendingMachine.requestItem()
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = vendingMachine.insertMoney(10)
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = vendingMachine.dispenseItem()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
