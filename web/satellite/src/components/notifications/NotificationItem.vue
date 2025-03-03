// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div :style="notification.style" class="notification-wrap" :class="{ active: isClassActive }" @mouseover="onMouseOver" @mouseleave="onMouseLeave">
        <div class="notification-wrap__content-area">
            <div class="notification-wrap__content-area__image">
                <component :is="notification.icon" />
            </div>
            <div class="notification-wrap__content-area__message-area">
                <p class="notification-wrap__content-area__message">{{ notification.message }}</p>
                <a
                    v-if="isSupportLinkMentioned"
                    :href="requestUrl"
                    class="notification-wrap__content-area__link"
                    target="_blank"
                >
                    {{ requestUrl }}
                </a>
            </div>
        </div>
        <div class="notification-wrap__buttons-group" @click="onCloseClick">
            <span class="notification-wrap__buttons-group__close">
                <CloseIcon />
            </span>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

import CloseIcon from '@/../static/images/notifications/close.svg';

import { DelayedNotification } from '@/types/DelayedNotification';
import { NOTIFICATION_ACTIONS } from '@/utils/constants/actionNames';
import { MetaUtils } from "@/utils/meta";

// @vue/component
@Component({
    components: {
        CloseIcon,
    },
})
export default class NotificationItem extends Vue {
    @Prop({default: () => new DelayedNotification(() => { return; }, '', '')})
    private notification: DelayedNotification;

    public readonly requestUrl = MetaUtils.getMetaContent('general-request-url');
    public isClassActive = false;

    /**
     * Indicates if support word is mentioned in message.
     * Temporal solution, can be changed later.
     */
    public get isSupportLinkMentioned(): boolean {
        return this.notification.message.toLowerCase().includes('support');
    }

    /**
     * Forces notification deletion.
     */
    public onCloseClick(): void {
        this.$store.dispatch(NOTIFICATION_ACTIONS.DELETE, this.notification.id);
    }

    /**
     * Forces notification to stay on page on mouse over it.
     */
    public onMouseOver(): void {
        this.$store.dispatch(NOTIFICATION_ACTIONS.PAUSE, this.notification.id);
    }

    /**
     * Resume notification flow when mouse leaves notification.
     */
    public onMouseLeave(): void {
        this.$store.dispatch(NOTIFICATION_ACTIONS.RESUME, this.notification.id);
    }

    /**
     * Uses for class change for animation.
     */
    public mounted(): void {
        setTimeout(() => {
            this.isClassActive = true;
        }, 100);
    }
}
</script>

<style scoped lang="scss">
    .notification-wrap {
        position: relative;
        right: -100%;
        width: calc(100% - 40px);
        height: auto;
        display: flex;
        justify-content: space-between;
        padding: 20px;
        align-items: center;
        border-radius: 12px;
        margin-bottom: 7px;
        transition: all 0.3s;

        &__content-area {
            display: flex;
            align-items: center;
            font-family: 'font_medium', sans-serif;
            font-size: 14px;

            &__image {
                max-height: 40px;
            }

            &__message-area {
                display: flex;
                flex-direction: column;
                align-items: flex-start;
                justify-content: space-between;
                margin: 0 0 0 17px;
            }

            &__message,
            &__link {
                height: auto;
                width: 270px;
                word-break: break-word;
            }

            &__link {
                margin-top: 5px;
                color: #224ca5;
                text-decoration: underline;
                cursor: pointer;
                word-break: normal;
            }
        }

        &__buttons-group {
            display: flex;

            &__close {
                width: 32px;
                height: 32px;
                cursor: pointer;
            }
        }
    }

    .active {
        right: 0;
    }
</style>
