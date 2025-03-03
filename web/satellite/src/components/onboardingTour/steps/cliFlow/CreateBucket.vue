// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <CLIFlowContainer
        :on-back-click="onBackClick"
        :on-next-click="onNextClick"
        title="Create a bucket"
    >
        <template #icon>
            <Icon />
        </template>
        <template #content class="create-bucket">
            <p class="create-bucket__msg">
                Let's create a bucket to store your data.<br><br>
                You can name your bucket using only lowercase alphanumeric characters (no spaces), like “cakes”.
            </p>
            <OSContainer>
                <template #windows>
                    <TabWithCopy value="./uplink.exe mb sj://cakes" aria-role-description="windows-create-bucket" />
                </template>
                <template #linux>
                    <TabWithCopy value="uplink mb sj://cakes" aria-role-description="linux-create-bucket" />
                </template>
                <template #macos>
                    <TabWithCopy value="uplink mb sj://cakes" aria-role-description="macos-create-bucket" />
                </template>
            </OSContainer>
        </template>
    </CLIFlowContainer>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import { RouteConfig } from "@/router";

import CLIFlowContainer from "@/components/onboardingTour/steps/common/CLIFlowContainer.vue";
import OSContainer from "@/components/onboardingTour/steps/common/OSContainer.vue";
import TabWithCopy from "@/components/onboardingTour/steps/common/TabWithCopy.vue";

import Icon from "@/../static/images/onboardingTour/bucketStep.svg";

import { AnalyticsHttpApi } from "@/api/analytics";

// @vue/component
@Component({
    components: {
        CLIFlowContainer,
        Icon,
        OSContainer,
        TabWithCopy,
    }
})
export default class CreateBucket extends Vue {

    private readonly analytics: AnalyticsHttpApi = new AnalyticsHttpApi();

    /**
     * Holds on back button click logic.
     */
    public async onBackClick(): Promise<void> {
        this.analytics.pageVisit(RouteConfig.OnboardingTour.with(RouteConfig.OnbCLIStep.with(RouteConfig.CLISetup)).path);
        await this.$router.push(RouteConfig.OnboardingTour.with(RouteConfig.OnbCLIStep.with(RouteConfig.CLISetup)).path);
    }

    /**
     * Holds on next button click logic.
     */
    public async onNextClick(): Promise<void> {
        this.analytics.pageVisit(RouteConfig.OnboardingTour.with(RouteConfig.OnbCLIStep.with(RouteConfig.UploadObject)).path);
        await this.$router.push(RouteConfig.OnboardingTour.with(RouteConfig.OnbCLIStep.with(RouteConfig.UploadObject)).path);
    }
}
</script>

<style scoped lang="scss">
    .create-bucket {
        font-family: 'font_regular', sans-serif;

        &__msg {
            font-size: 16px;
            line-height: 24px;
            color: #1b2533;
            margin-bottom: 20px;
        }
    }
</style>
