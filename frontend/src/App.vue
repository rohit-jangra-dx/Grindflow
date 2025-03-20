<script setup lang="ts">
import SideBar from "./components/sidebar/SideBar.vue";
import {
    NConfigProvider,
    NLayout,
    NLayoutContent,
    NLayoutFooter,
    NMessageProvider,
    NNotificationProvider,
    NModalProvider
} from "naive-ui";
import Notification from "./components/Notification.vue";


// theme store
import { useThemeStore } from "./stores/theme";
import MainHeader from "./components/header/MainHeader.vue";

const themeStore = useThemeStore()
</script>

<template>
    <n-config-provider :theme="themeStore.getTheme">
        <!-- message area -->
        <NMessageProvider>
            <!-- Dialog area -->
            <NModalProvider>
                <!-- Notification area -->
                <NNotificationProvider>
                    <NLayout has-sider class="app-layout">
                        <SideBar />

                        <!-- another layout wrapper -->
                        <NLayout>
                            <MainHeader/>

                            <!-- content goes here acts as a page -->
                            <NLayoutContent>
                                 
                                <RouterView/>

                            </NLayoutContent>

                            <NLayoutFooter>
                                <Notification />
                            </NLayoutFooter>
                        </NLayout>
                    </NLayout>
                </NNotificationProvider>
            </NModalProvider>
        </NMessageProvider>
    </n-config-provider>
</template>

<style scoped>
.app-layout {
    height: 100vh;
}
</style>
