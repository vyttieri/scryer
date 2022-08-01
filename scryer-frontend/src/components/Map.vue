<script setup>
import { storeToRefs } from 'pinia'

import { useDeviceStore } from '@/stores/device'
import { useDevicePreferencesStore } from '@/stores/devicePreferences'

const deviceStore = useDeviceStore()
const { getDeviceIcon } = useDevicePreferencesStore()
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
			@click="$emit('set-center', device.latest_accurate_device_point.device_point_detail.lat_lng)"
			:label="{ className: 'q-icons material-icons', fontFamily: 'Material Icons', text: getDeviceIcon(device), fontSize: '18px', color: '#fff' }"
		/>
	</GMapMap>
</template>

<script>
export default {
	props: {
		center: { lat: String, lng: String, },
		setCenter: Function,
	}
}
</script>
