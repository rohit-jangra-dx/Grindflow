<script setup lang="ts">
import { NButton, NFormItem, NInput, NModal, useMessage } from "naive-ui";
import { CreateUser } from "../../wailsjs/go/main/App";
import { ref } from "vue";

// toggle for modal
const showModal = ref<boolean>(false);
const userName = ref<string>("");

// for showing messages
const message = useMessage();

const createNewUser = async (username: string) => {
    try {
        const data = await CreateUser(username);
        message.success(data);
    } catch (err: unknown) {
        //when something went wrong
        if (err instanceof Error) {
            message.error(err.message);
            return;
        }
        message.error("Couldn't catch clear error js is fj");
    }
};

const cancelCallback = () => {
    message.warning("You need to specify a username to continue");
};
</script>

<template>
    <div>
        <NButton @click="showModal = true"> Create a New User</NButton>
        <NModal
            v-model:show="showModal"
            preset="dialog"
            title="Need a username to continue!"
            positive-text="Continue"
            negative-text="Cancel"
            @positive-click="createNewUser(userName)"
            @negative-click="cancelCallback"
        >
            
            <NFormItem style="margin-top: 1rem;" label="Your Username">
                <NInput
                    v-model:value="userName"
                    type="text"
                    placeholder="e.g. rohit_232!"
                />
            </NFormItem>
        </NModal>
    </div>
</template>
