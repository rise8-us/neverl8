<template>
  <v-container class="d-flex justify-end">
    <v-date-picker
      show-adjacent-months
      :model-value="date"
      :min="minDate"
      :max="maxDate"
      @update:model-value="handleDateInput"
    ></v-date-picker>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      date: null,
    };
  },
  computed: {
    minDate() {
      const today = new Date();
      today.setDate(today.getDate() - 1);
      return today;
    },
    maxDate() {
      const today = new Date();
      today.setDate(today.getDate() + 30);
      return today;
    },
  },
  methods: {
    formatDate(dateString) {
      const date = new Date(dateString);
      const year = date.getFullYear();
      const month = (date.getMonth() + 1).toString().padStart(2, "0");
      const day = date.getDate().toString().padStart(2, "0");

      return `${year}-${month}-${day}`;
    },
    handleDateInput(date) {
      this.date = date;
      this.$emit("dateSelected", this.formatDate(date));
    },
  },
};
</script>
