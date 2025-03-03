// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <section class="payout-history-table">
        <div class="payout-history-table__header">
            <p class="payout-history-table__header__title">Payout History</p>
            <div class="payout-history-table__header__selection-area">
                <PayoutHistoryPeriodDropdown />
            </div>
        </div>
        <div class="payout-history-table__divider" />
        <div class="payout-history-table__table-container">
            <div class="payout-history-table__table-container__labels-area">
                <p class="payout-history-table__table-container__labels-area__label">Satellite</p>
                <p class="payout-history-table__table-container__labels-area__label">Payout</p>
            </div>
            <PayoutHistoryTableItem v-for="historyItem in payoutHistory" :key="historyItem.satelliteID" :history-item="historyItem" />
            <div class="payout-history-table__table-container__totals-area">
                <p class="payout-history-table__table-container__totals-area__label">Total</p>
                <p class="payout-history-table__table-container__totals-area__value">{{ totalPaid | centsToDollars }}</p>
            </div>
        </div>
    </section>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import PayoutHistoryPeriodDropdown from '@/app/components/payments/PayoutHistoryPeriodDropdown.vue';
import PayoutHistoryTableItem from '@/app/components/payments/PayoutHistoryTableItem.vue';

import { PAYOUT_ACTIONS } from '@/app/store/modules/payout';
import { SatellitePayoutForPeriod } from '@/storagenode/payouts/payouts';

// @vue/component
@Component ({
    components: {
        PayoutHistoryPeriodDropdown,
        PayoutHistoryTableItem,
    },
})
export default class PayoutHistoryTable extends Vue {
    public get payoutHistory(): SatellitePayoutForPeriod[] {
        return this.$store.state.payoutModule.payoutHistory;
    }

    /**
     * Returns sum of payouts for all satellites of current period.
     */
    public get totalPaid(): number {
        return this.$store.getters.totalPaidForPayoutHistoryPeriod;
    }

    /**
     * Lifecycle hook after initial render.
     * Fetches payout history for last period.
     */
    public async mounted(): Promise<void> {
        const payoutPeriods = this.$store.state.payoutModule.payoutPeriods;

        if (!payoutPeriods.length) {
            return;
        }

        const lastPeriod = payoutPeriods[payoutPeriods.length - 1];
        await this.$store.dispatch(PAYOUT_ACTIONS.SET_PAYOUT_HISTORY_PERIOD, lastPeriod.period);

        try {
            await this.$store.dispatch(PAYOUT_ACTIONS.GET_PAYOUT_HISTORY);
        } catch (error) {
            console.error(error);
        }
    }
}
</script>

<style scoped lang="scss">
    .payout-history-table {
        display: flex;
        flex-direction: column;
        padding: 20px 40px;
        background: var(--block-background-color);
        border: 1px solid var(--block-border-color);
        box-sizing: border-box;
        border-radius: 12px;
        font-family: 'font_regular', sans-serif;

        &__header {
            display: flex;
            flex-direction: row;
            align-items: center;
            justify-content: space-between;
            height: 40px;

            &__title {
                font-weight: 500;
                font-size: 18px;
                color: var(--regular-text-color);
            }
        }

        &__divider {
            width: 100%;
            height: 1px;
            background-color: #eaeaea;
        }

        &__table-container {
            width: 100%;
            margin-top: 19px;

            &__labels-area,
            &__totals-area {
                display: flex;
                align-items: center;
                justify-content: space-between;
                padding: 8px 17px;
                width: calc(100% - 34px);
                font-family: 'font_medium', sans-serif;
            }

            &__labels-area {
                background: var(--table-header-color);

                &__label {
                    font-size: 14px;
                    color: var(--label-text-color);
                }
            }

            &__totals-area {
                margin-top: 12px;
                font-size: 16px;
                color: var(--regular-text-color);
            }
        }
    }

    @media screen and (max-width: 640px) {

        .payout-history-table {
            padding: 28px 20px;

            &__divider {
                display: none;
            }
        }
    }
</style>
