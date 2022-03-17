package main

import "github.com/Kininaru/operating-system-practicum/utils"

func fifo(memory []int, data int) {
	for i := 0; i < len(memory)-1; i++ {
		memory[i] = memory[i+1]
	}
	memory[len(memory)-1] = data
}

func FirstInFirstOut(page int, queue []int) {
	memory := make([]int, page)
	for index := range memory {
		memory[index] = -1
	}

	for _, data := range queue {
		index := utils.GetIndex(memory, data)
		if index == -1 {
			fifo(memory, data)
		}
		utils.PrintMemory(memory)
	}
}