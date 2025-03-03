// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <VModal :on-close="closeModal">
        <template #content>
            <div class="modal">
                <img
                    class="modal__icon"
                    src="@/../static/images/account/billing/paidTier/prompt.png"
                    alt="Prompt Image"
                >
                <h1 class="modal__title" aria-roledescription="modal-title">
                    Get more projects<br>when you upgrade
                </h1>
                <p class="modal__info">
                    Upgrade your Free Account to create<br>more projects and gain access to higher limits.
                </p>
                <VButton
                    width="256px"
                    height="56px"
                    border-radius="8px"
                    label="Upgrade to Pro Account ->"
                    :on-press="onClick"
                />
            </div>
        </template>
    </VModal>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import { APP_STATE_MUTATIONS } from "@/store/mutationConstants";

import VButton from '@/components/common/VButton.vue';
import VModal from '@/components/common/VModal.vue';

// @vue/component
@Component({
    components: {
        VButton,
        VModal,
    },
})
export default class CreateProjectPromptModal extends Vue {
    /**
     * Holds on button click logic.
     * Closes this modal and opens upgrade account modal.
     */
    public onClick(): void {
        this.$store.commit(APP_STATE_MUTATIONS.TOGGLE_CREATE_PROJECT_PROMPT_POPUP);
        this.$store.commit(APP_STATE_MUTATIONS.TOGGLE_IS_ADD_PM_MODAL_SHOWN);
    }

    /**
     * Closes create project prompt modal.
     */
    public closeModal(): void {
        this.$store.commit(APP_STATE_MUTATIONS.TOGGLE_CREATE_PROJECT_PROMPT_POPUP);
    }
}
</script>

<style scoped lang="scss">
    .modal {
        width: 600px;
        padding: 50px 0 65px;
        display: flex;
        align-items: center;
        flex-direction: column;
        font-family: 'font_regular', sans-serif;

        &__icon {
            max-height: 154px;
            max-width: 118px;
        }

        &__title {
            font-family: 'font_bold', sans-serif;
            font-size: 28px;
            line-height: 34px;
            color: #1b2533;
            margin-top: 40px;
            text-align: center;
        }

        &__info {
            font-family: 'font_regular', sans-serif;
            font-size: 16px;
            line-height: 21px;
            text-align: center;
            color: #354049;
            margin: 15px 0 45px;
        }
    }
</style>
