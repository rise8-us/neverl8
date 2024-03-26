<template>
  <v-app>
    <v-container>
      <v-row class="d-flex justify-center">
        <h1>Neverl8 Meeting Scheduler</h1></v-row
      >
      <v-divider class="my-4"></v-divider>
      <v-row>
        <v-col>
          <CalendarView @dateSelected="handleDateSelected" />
        </v-col>
        <v-col>
          <ScheduleMeeting
            v-if="meeting"
            :meeting="meeting"
            :availableTimeSlots="availableTimeSlots"
            @scheduleMeeting="handleScheduleMeeting"
          />
        </v-col>
      </v-row>
      <v-button @click="createMeeting">Create Meeting (test)</v-button>
    </v-container>
  </v-app>
</template>

<script>
import CalendarView from "./components/CalendarView.vue";
import ScheduleMeeting from "./components/ScheduleMeeting.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    CalendarView,
    ScheduleMeeting,
  },
  data() {
    return {
      meeting: null,
      availableTimeSlots: [],
    };
  },

  mounted() {
    this.fetchMeetingDetails();
  },
  methods: {
    async createMeeting() {
      try {
        await axios.post("/api/meeting", {
          title: "Test Meeting",
          duration: 60,
          description: "This is a test meeting",
          calendar: "test",
          Hosts: [
            {
              host_name: "Test Host",
              time_preferences: [
                {
                  id: 1,
                  host_id: 1,
                  start_time: "2024-03-26T09:00:00Z",
                  end_time: "2024-03-26T20:00:00Z",
                },
              ],
            },
          ],
        });
      } catch (error) {
        console.error("Error creating meeting:", error);
      }
    },
    async fetchMeetingDetails() {
      // Fetch meeting details based on hardcoded ID
      try {
        const response = await axios.get(`/api/meeting/1`);
        this.meeting = response.data;
      } catch (error) {
        console.error("Error fetching meeting:", error);
      }
    },
    async handleDateSelected(date) {
      this.selectedTimeSlotID = -1;
      this.selectedSlot = null;

      try {
        const response = await axios.get(
          `/api/meeting/time-slots/1?date=${date}` // TODO: Replace hardcoded ID with dynamic ID based on URL sent to candidate
        );
        this.availableTimeSlots = response.data;
      } catch (error) {
        console.error("Error fetching available time slots:", error);
        this.availableTimeSlots = [];
      }
    },
    async handleScheduleMeeting(
      selectedTimeSlot,
      candidateName,
      candidateEmail
    ) {
      try {
        await axios.post(`/api/meeting/schedule`, {
          meeting_id: this.meeting.id,
          time_slot: selectedTimeSlot,
          candidate_name: candidateName,
          candidate_email: candidateEmail,
        });
        alert("Meeting scheduled!");
      } catch (error) {
        console.error("Error scheduling meeting:", error);
        alert("Error scheduling meeting");
      }
    },
  },
};
</script>
