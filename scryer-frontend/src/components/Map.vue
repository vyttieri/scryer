<script setup>
	import { storeToRefs } from 'pinia'

	import LoginForm from './LoginForm.vue'
	import RegistrationForm from './RegistrationForm.vue'

	import { useDeviceStore } from '@/stores/device'
	import { useDevicePreferencesStore } from '@/stores/devicePreferences'

	const { devices, visibleDevices } =  storeToRefs(useDeviceStore())

	const californiaCenter = { lat: 36.778300, lng: -119.417900 }

</script>

<!-- TODO: Centering seems wrong; device_point_detail is null -->
<template>
	<LoginForm />
	<RegistrationForm />
	<GMapMap
		:center="californiaCenter"
		:zoom="6"
	>
		<GMapMarker
			:key="device.device_id + '-marker'"
			:position="device.latest_accurate_device_point.device_point_detail.lat_lng"
			v-for="device in visibleDevices"
		/>
	</GMapMap>
</template>
