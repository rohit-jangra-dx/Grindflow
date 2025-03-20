<script setup lang="ts">
import { NLayoutSider} from 'naive-ui'
import MainMenu from './menu/MainMenu.vue';
import { ref } from 'vue';

const collapsed = ref<boolean>(false)

// controls collapsed via this
const toggleCollapsed= () => { collapsed.value = !collapsed.value}

</script>

<template>
    <NLayoutSider 
    bordered
    show-trigger
    :collapsed="collapsed"
    collapse-mode="width"
    :collapsed-width="64"
    :width="240"
    :native-scrollbar="false"
    :inverted="false"
    @collapse="toggleCollapsed"
    @expand="toggleCollapsed"
    >
    
    <MainMenu :collapsed="collapsed"/>
    </NLayoutSider>
</template>