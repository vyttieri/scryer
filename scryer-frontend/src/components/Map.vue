<script setup>
import { storeToRefs } from 'pinia'
import { defineProps } from 'vue'

import { useDeviceStore } from '@/stores/device'
import { useDevicePreferenceStore } from '@/stores/devicePreferences'

defineProps({
	center: { type: Object, required: true },
	openedMarkerId: { type: String, required: true },
})

const deviceStore = useDeviceStore()
const { getDeviceIcon } = useDevicePreferenceStore()
const { visibleDevices } =  storeToRefs(deviceStore)
</script>

<template>
	<GMapMap
		:center="center"
		:zoom="6"
	>
		<GMapMarker
			:key="device.device_id + '-marker'"
			v-for="device in visibleDevices"
			:position="device.latest_accurate_device_point.device_point_detail.lat_lng"
			@click="$emit('set-open-marker', device.device_id)"
			:label="{ className: 'q-icons material-icons', fontFamily: 'Material Icons', text: getDeviceIcon(device), fontSize: '18px', color: '#fff' }"
		>
			<GMapInfoWindow :opened="openedMarkerId == device.device_id">
				{{ device.display_name}}
			</GMapInfoWindow>
		</GMapMarker>
	</GMapMap>
</template>
