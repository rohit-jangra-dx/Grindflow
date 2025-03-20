<script setup lang="ts">
import { ref } from "vue";
import SideBar from "./components/sidebar/SideBar.vue";
import Signup from "./components/Signup.vue";
import {
    NConfigProvider,
    NLayout,
    NLayoutContent,
    NLayoutFooter,
    NLayoutHeader,
    NMessageProvider,
    NNotificationProvider,
    NModalProvider
} from "naive-ui";
import Notification from "./components/Notification.vue";

import ThemeSwitch from "./components/ThemeSwitch.vue";

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
                            
                            <NLayoutContent> <Signup /> </NLayoutContent>

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
