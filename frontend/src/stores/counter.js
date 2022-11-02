import { ref, computed } from 'vue'
import { defineStore } from 'pinia'


// TODO: Look into Pinia and see if we want to keep it in the project
export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }

  return { count, doubleCount, increment }
})
