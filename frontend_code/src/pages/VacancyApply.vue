<template>
  <div>
    <Header />
    <main class="vacancy-apply-form">
      <h1 class="vacancy-apply-form__title">Отклик на вакансию</h1>
      <form @submit.prevent="submitForm" enctype="multipart/form-data">
        <input v-model="form.name" type="text" placeholder="Имя *" />
        <input v-model="form.surname" type="text" placeholder="Фамилия *" />
        <input
          v-model="form.lastname"
          type="text"
          placeholder="Отчество * (если нет — введите прочерк)"
        />
        <input
          v-model="form.phone"
          type="tel"
          placeholder="Телефон * (например, 89995556677)"
        />

        <label class="resume-label">Прикрепите ваше резюме (PDF) *</label>
        <input type="file" @change="handleFile" accept=".pdf" />

        <button type="submit">Отправить заявку</button>
      </form>

      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </main>
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const form = ref({
  name: '',
  surname: '',
  lastname: '',
  phone: '',
  file: null
})

const errorMessage = ref('')

function handleFile(event) {
  form.value.file = event.target.files[0]
}

function isValidPhone(phone) {
  return /^\d{11}$/.test(phone)
}

async function submitForm() {
  const { name, surname, lastname, phone, file } = form.value

  if (!name || !surname || !lastname || !phone || !file) {
    errorMessage.value = 'Пожалуйста, заполните все поля и прикрепите файл резюме в формате PDF.'
    return
  }

  if (!isValidPhone(phone)) {
    errorMessage.value = 'Введите корректный номер телефона (11 цифр, например 89995556677).'
    return
  }

  if (file.type !== 'application/pdf') {
    errorMessage.value = 'Резюме должно быть в формате PDF.'
    return
  }

  errorMessage.value = '' // очистка ошибки

  const fd = new FormData()
  fd.append('name', name)
  fd.append('surname', surname)
  fd.append('lastname', lastname)
  fd.append('phone', phone)
  fd.append('vacancy_id', route.params.id)
  fd.append('resume', file)

  const res = await fetch('/api/vacancy/apply', {
    method: 'POST',
    body: fd
  })

  if (res.ok) {
    alert('Заявка отправлена!')
  } else {
    alert('Ошибка при отправке.')
  }
}
</script>

<style scoped lang="sass">
.vacancy-apply-form
  max-width: 600px
  margin: 0 auto
  padding: 0px 10px 30px 10px
  &__title
    color: #1C5242
    font-weight: bold
    font-size: 2rem

form
  display: flex
  flex-direction: column
  gap: 1rem

input, button
  padding: 0.75rem
  font-size: 1rem

.resume-label
  font-weight: bold
  margin-top: 1rem

button
  background-color: #1C5242
  color: white
  border: none
  border-radius: 6px
  cursor: pointer
  transition: background-color 0.3s

button:hover
  background-color: #218838

.error-message
  color: red
  margin-top: 1rem
  font-weight: bold
</style>
