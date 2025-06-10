<template>
    <div class="person-card">
        <img
      class="person-card__image"
      :src="getImageUrl(person.photo)"
      alt="Фото сотрудника"
    />
    <div class="person-card__text">
      <p><b>{{ person.surname}} {{person.name}} {{person.lastname}} </b></p>
      <p>{{ person.vacancy }} </p>
      <p>{{ person.phone_number }} </p>
    </div>
    </div>
</template>

<script setup>
const props = defineProps({
  person: {
    type: Object,
    required: true
  }
});

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg';
}

</script>

<style lang="sass" scoped>
@import '@/assets/style.sass'
</style>