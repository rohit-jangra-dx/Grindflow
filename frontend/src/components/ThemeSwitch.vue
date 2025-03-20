<script setup lang="ts">
import { NSwitch, NIcon } from 'naive-ui';
import {WbSunnyFilled, DarkModeFilled} from '@vicons/material'
import { CSSProperties, ref } from 'vue';
import { useThemeStore } from '../stores/theme';

const themeStore = useThemeStore()

// for switch state
const IsSwitchOn = ref<boolean>(themeStore.getTheme.name == 'light')

type railStyleArgs = {
    focused: boolean;
    checked: boolean;
}
// switch style
const railStyle = ( {focused, checked} : railStyleArgs) => {
    const style: CSSProperties = {}
    if (checked) {
        style.background = '#000000'
        if (focused) {
            style.boxShadow = '0 0 0 2px #00000000'
        }
    } else {
        style.background = '#2080f0'
        if (focused) {
            style.boxShadow = '0 0 0 2px #2080f040'
        }
    }
    return style
}


</script>

<template>
    <NSwitch 
    :rail-style="railStyle" 
    v-model:value="IsSwitchOn"
    @update:value="themeStore.toggleTheme()"
    >
        <template #checked-icon>
            <NIcon :component="WbSunnyFilled"/>
        </template>
        <template #unchecked-icon>
            <NIcon :component="DarkModeFilled"/>
            
            
        </template>
    </NSwitch>
</template>