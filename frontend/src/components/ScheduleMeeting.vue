<template>
  <div>
    <h2>Schedule Meeting</h2>
    <div v-if="meeting">
      <p>Meeting: {{ meeting.title }}</p>
      <p>Duration: {{ meeting.duration }} minutes</p>
      <div v-for="host in meeting.Hosts" :key="host.id">
        <ul v-if="isHostInAvailableSlots(host.id)">
          {{
            host.host_name
          }}
          is available for a meeting at the following times
        </ul>
      </div>
      <div v-if="availableTimeSlots.length > 0">
        <h3>Available Time Slots:</h3>
        <ul>
          <li v-for="slot in availableTimeSlots" :key="slot.id">
            {{ formatDate(slot.start_time) }} -
            {{ formatDate(slot.end_time) }}
            <button
              @click="selectTimePreference(slot), console.log(slot)"
              :class="{
                'selected-time-slot': selectedTimeSlotID === slot.id,
              }"
            >
              Select
            </button>
          </li>
        </ul>
      </div>
      <div v-else>
        <h3>No available time slots for this day</h3>
      </div>
      <div v-if="selectedTimeSlotID >= 0">
        <h4>Enter your details</h4>
        <form @submit.prevent="scheduleMeeting">
          <input
            type="text"
            v-model="candidateName"
            placeholder="Your Name"
            required
          />
          <input
            type="email"
            v-model="candidateEmail"
            placeholder="Your Email"
            required
          />
          <button type="submit">Submit</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    meeting: Object,
    availableTimeSlots: Array,
  },
  data() {
    return {
      candidateName: "",
      candidateEmail: "",
      selectedTimeSlotID: -1,
    };
  },
  methods: {
    isHostInAvailableSlots(hostId) {
      return this.availableTimeSlots.some((slot) => slot.host_id === hostId);
    },
    selectTimePreference(selectedSlot) {
      this.selectedTimeSlotID = selectedSlot.id;
    },
    scheduleMeeting() {
      this.$emit("scheduleConfirmed");
      // Reset form
      this.candidateName = "";
      this.candidateEmail = "";
      this.selectedTimeSlotID = -1;
    },
    formatDate(dateString) {
      const date = new Date(dateString);
      return date.toLocaleTimeString("en-US", {
        hour: "numeric",
        minute: "numeric",
      });
    },
  },
};
</script>

<style>
.selected-time-slot {
  background-color: #4caf50;
  color: white;
}
</style>
