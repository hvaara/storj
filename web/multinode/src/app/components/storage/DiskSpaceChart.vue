// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="chart">
        <p class="disk-space-chart__data-dimension">{{ chartDataDimension }}*h</p>
        <VChart
            id="disk-space-chart"
            :key="chartKey"
            :chart-data="chartData"
            :width="chartWidth"
            :height="chartHeight"
            :tooltip-constructor="diskSpaceTooltip"
        />
    </div>
</template>

<script lang="ts">
import { Component } from 'vue-property-decorator';

import BaseChart from '@/app/components/common/BaseChart.vue';
import VChart from '@/app/components/common/VChart.vue';

import { ChartData, Tooltip, TooltipParams, TooltipModel } from '@/app/types/chart';
import { Chart as ChartUtils } from '@/app/utils/chart';
import { Size } from '@/private/memory/size';
import { Stamp } from '@/storage';

/**
 * stores stamp data for disc space chart's tooltip
 */
class StampTooltip {
    public atRestTotal: string;
    public date: string;

    public constructor(stamp: Stamp) {
        this.atRestTotal = Size.toBase10String(stamp.atRestTotal);
        this.date = stamp.intervalStart.toUTCString().slice(0, 16);
    }
}

// @vue/component
@Component({
    components: {VChart},
})
export default class DiskSpaceChart extends BaseChart {
    private get allStamps(): Stamp[] {
        return ChartUtils.populateEmptyStamps(this.$store.state.storage.usage.diskSpaceDaily);
    }

    public get chartDataDimension(): string {
        if (!this.$store.state.storage.usage.diskSpaceDaily.length) {
            return 'Bytes';
        }

        return ChartUtils.getChartDataDimension(this.allStamps.map((elem) => elem.atRestTotal));
    }

    public get chartData(): ChartData {
        let data: number[] = [0];
        const daysCount = ChartUtils.daysDisplayedOnChart();
        const chartBackgroundColor = '#4F97F7';
        const chartBorderColor = '#1F49A3';
        const chartBorderWidth = 1;

        if (this.allStamps.length) {
            data = ChartUtils.normalizeChartData(this.allStamps.map(elem => elem.atRestTotal));
        }

        return new ChartData(daysCount, chartBackgroundColor, chartBorderColor, chartBorderWidth, data);
    }

    public diskSpaceTooltip(tooltipModel: TooltipModel): void {
        const tooltipParams = new TooltipParams(tooltipModel, 'disk-space-chart', 'disk-space-tooltip', 'disk-space-tooltip-point', this.tooltipMarkUp(tooltipModel),
            125, 89, 6, 4, '#1f49a3');

        Tooltip.custom(tooltipParams);
    }

    private tooltipMarkUp(tooltipModel: TooltipModel): string {
        if (!tooltipModel.dataPoints) {
            return '';
        }

        const dataIndex = tooltipModel.dataPoints[0].index;
        const dataPoint = new StampTooltip(this.allStamps[dataIndex]);

        return `<div class='tooltip-body'>
                    <p class='tooltip-body__data'><b>${dataPoint.atRestTotal}*h</b></p>
                    <p class='tooltip-body__footer'>${dataPoint.date}</p>
                </div>`;
    }
}
</script>

<style lang="scss">
    p {
        margin: 0;
    }

    .disk-space-chart {
        z-index: 102;

        &__data-dimension {
            font-size: 13px;
            color: var(--c-title);
            margin: 0 0 5px 31px !important;
            font-family: 'font_medium', sans-serif;
        }
    }

    #disk-space-tooltip {
        background: white;
        border: 1px solid var(--c-gray--light);
        width: 180px;
        height: 90px;
        font-size: 12px;
        border-radius: 14px;
        color: var(--c-title);
        pointer-events: none;
        z-index: 9999;
        font-family: 'font_bold', sans-serif;
    }

    .tooltip-body {

        &__data {
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 11px 44px;
            font-size: 14px;
        }

        &__footer {
            font-size: 12px;
            width: auto;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 10px 0;
        }
    }
</style>
