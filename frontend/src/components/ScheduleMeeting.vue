<template>
  <v-container>
    <v-row v-if="meeting" justify="start">
      <v-col cols="12" sm="8" md="8" lg="8" xl="6" xxl="5">
        <v-card outlined>
          <v-card-title> Meeting: {{ meeting.title }} </v-card-title>
          <v-card-text>
            Duration: {{ meeting.duration }} minutes
            <v-divider class="my-3"></v-divider>
            <div v-for="host in meeting.Hosts" :key="host.id">
              <v-chip v-if="isHostInAvailableSlots(host.id)" class="mb-2">
                {{ host.host_name }} is available for meetings
              </v-chip>
            </div>
            <v-divider
              class="my-3"
              v-if="availableTimeSlots.length > 0"
            ></v-divider>
            <v-subheader>Available Time Slots:</v-subheader>
            <v-virtual-scroll
              :items="availableTimeSlots"
              max-height="270"
              v-if="availableTimeSlots.length > 0"
            >
              <template v-slot:default="{ item }">
                <v-list-item
                  @click="selectTimePreference(item)"
                  :active="item.id === selectedTimeSlotID"
                >
                  {{ formatDate(item.start_time) }} -
                  {{ formatDate(item.end_time) }}
                </v-list-item>
              </template>
            </v-virtual-scroll>
            <div v-else>
              <v-alert border="left" type="info" col="12"
                >No available time slots for this day</v-alert
              >
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="6" v-if="selectedTimeSlotID >= 0">
        <v-card outlined>
          <v-card-title>Enter your details</v-card-title>
          <v-card-text>
            <v-form @submit.prevent="scheduleMeeting">
              <v-text-field
                v-model="candidateName"
                label="Your Name"
                required
              ></v-text-field>
              <v-text-field
                v-model="candidateEmail"
                label="Your Email"
                type="email"
                required
              ></v-text-field>
              <v-btn type="submit" color="primary">Submit</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
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
      selectedSlot: null,
    };
  },
  methods: {
    isHostInAvailableSlots(hostId) {
      return this.availableTimeSlots.some((slot) => slot.host_id === hostId);
    },
    selectTimePreference(selectedSlot) {
      this.selectedTimeSlotID = selectedSlot.id;
      this.selectedSlot = selectedSlot;
    },
    scheduleMeeting() {
      this.$emit(
        "scheduleMeeting",
        this.selectedSlot,
        this.candidateName,
        this.candidateEmail
      );
      // Reset form
      this.candidateName = "";
      this.candidateEmail = "";
      this.selectedTimeSlotID = -1;
      this.selectedSlot = null;
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
