<script setup lang="ts">

import { MenuOption, NButton, NIcon, NMenu } from 'naive-ui';
import { h, ref } from 'vue';
import { DashboardCustomizeOutlined, PendingActionsFilled, ScheduleFilled, WorkOutlineRound, AddRound } from '@vicons/material'
import { Dashboard, TagGroup } from '@vicons/carbon'
import { RouterLink } from 'vue-router';


const { collapsed } = defineProps<{ collapsed: boolean }>()

console.log(collapsed)

// Menu options
// FIXME: correctly setup the menu options
const menuOptions: MenuOption[] = [
    {
        label: () => h(RouterLink, { to: '/', style: 'text-decoration:none, color: inherit;' }, 'Dashboard'),
        icon: () => h(NIcon, null, { default: () => h(Dashboard) }),
        key: "dashboard"
    },
    {
        label: () => h(RouterLink, { to: '/plan', style: 'text-decoration:none, color: inherit;' }, 'Plan'),
        key: "plan",
        icon: () => h(NIcon, null, { default: () => h(PendingActionsFilled) }),
    },
    {
        label: () => h(RouterLink, { to: '/schedule', style: 'text-decoration:none, color: inherit;' }, 'Schedule'),
        key: "schedule",
        icon: () => h(NIcon, null, { default: () => h(ScheduleFilled) }),
    },
    {
        label: () => h(RouterLink, { to: '/boards', style: 'text-decoration:none, color: inherit;' }, 'Boards'),
        key: "boards",
        icon: () => h(NIcon, null, { default: () => h(DashboardCustomizeOutlined) }),
    },
    {
        key: 'divider-1',
        type: 'divider',
        props: {
            style: {
                marginLeft: '32px',
                marginBottom: "16px"
            }
        }
    },
    {
        label: "Projects",
        key: "projects",
        icon: () => h(NIcon, null, { default: () => h(WorkOutlineRound) }),
        children: [
            {
                label: "Lodu",
                key: 'd'
            },
            {
                label: "chodu",
                key: 'v'
            }
        ],

    },
    {
        label: "Tags",
        key: "tags",
        icon: () => h(NIcon, null, { default: () => h(TagGroup) }),
        children: [
            {
                label: "Lodu",
                key: 'd'
            },
            {
                label: "chodu",
                key: 'v'
            }
        ]
    }
]



</script>

<template>
    <NMenu :collapsed="collapsed" style="padding-top: 10px;" :options="menuOptions" :collapsed-width="64"
        :collapsed-icon-size="22" />

</template>

<!-- wo ashman jhuk rha h zameen par -->