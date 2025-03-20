import {defineStore} from 'pinia'
import { darkTheme, GlobalTheme, lightTheme } from 'naive-ui'
import { computed, ref } from 'vue'


export const  useThemeStore = defineStore('theme', () => {
    const theme = ref<GlobalTheme>(lightTheme)

    // getters
    const getTheme = computed(() => theme.value)

    // actions
    // true for lighttheme false for dark
    const setTheme = (flag: boolean) => {
        theme.value = flag ? lightTheme : darkTheme
        return
    }
    
    // check if the light theme enabled or not i.e. it's default theme
    const toggleTheme = () => {
        const isDark = theme.value.name == 'light'
        setTheme(!isDark)
    }

    return { getTheme, setTheme, toggleTheme}
})