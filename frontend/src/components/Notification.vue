<script setup lang="ts">
import type { NotificationType } from 'naive-ui'
import { useNotification, NButton } from 'naive-ui'

const notification = useNotification()

const notify = (type: NotificationType) => {
  notification[type]({
    content: "What to say?",
    meta: "I don't know",
    duration: 2300,
    keepAliveOnHover: false
  })
}
</script>

<template>
    <NButton @click="notify('warning')"> Click me to see </NButton>
</template>