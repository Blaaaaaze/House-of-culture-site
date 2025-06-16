<template>
  <div class="page">
    <Header />
    <div class="page__content">
      <div class="group-single" v-if="group">
        <router-link class="news-card__link news-card__link_gr" to="/groups">Вернуться назад</router-link>

        <div class="group-single__content">
          <!-- Изображение группы -->
          <img
            :src="getImageUrl(groupImagePath)"
            alt="Изображение группы"
            class="group-single__image"
          />


          <div class="group-single__info">
            <h1 class="group-single__title">{{ group.name }}</h1>
            <p class="group-single__description" v-html="group.full_description"></p>

            <!-- Карточка преподавателя -->
            <div class="teacher-card" v-if="teacher">
              <img
                :src="getImageUrl(teacher.photo)"
                alt="Фото преподавателя"
                class="teacher-card__photo"
              />
              <div class="teacher-card__info">
                <h2 class="teacher-card__name">
                  {{ teacher.surname }} {{ teacher.name }} {{ teacher.lastname }}
                </h2>
                <p class="teacher-card__phone">Телефон: {{ teacher.phone_number }}</p>
              </div>
            </div>

            <!-- Кнопка и форма -->
            <button v-if="!showForm" class="apply-button" @click="showForm = true">
              Подать заявку
            </button>

            <form
              v-else
              class="apply-form"
              @submit.prevent="submitApplication"
            >
              <h2 class="apply-form__title">Заявка на вступление</h2>

              <input v-model="form.child_surname" type="text" placeholder="Фамилия ребенка *" />
              <input v-model="form.child_name" type="text" placeholder="Имя ребенка *" />
              <input v-model="form.child_lastname" type="text" placeholder="Отчество ребенка *" />

              <input v-model="form.parent_surname" type="text" placeholder="Фамилия родителя *" />
              <input v-model="form.parent_name" type="text" placeholder="Имя родителя *" />
              <input v-model="form.parent_lastname" type="text" placeholder="Отчество родителя *" />
              <input
                v-model="form.parent_phone"
                type="tel"
                placeholder="Телефон родителя * (например, 89995556677)"
              />

              <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>

              <button type="submit">Отправить</button>
              <button type="button" class="cancel-button" @click="resetForm">
                Отмена
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const group = ref(null)
const teacher = ref(null)
const groupImagePath = ref('')
const showForm = ref(false)
const errorMessage = ref('')

const form = ref({
  child_surname: '',
  child_name: '',
  child_lastname: '',
  parent_surname: '',
  parent_name: '',
  parent_lastname: '',
  parent_phone: ''
})

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg'
}

onMounted(async () => {
  try {
    const id = route.params.id
    const res = await fetch(`/api/groups/${id}`)
    const data = await res.json()

    group.value = data.group
    teacher.value = data.teacher

    // Получаем изображение группы из media
    const media = data.media || []
    const groupImage = media.find(
      m => m.type === 'image' && m.related_table === 'group'
    )
    if (groupImage) {
      groupImagePath.value = groupImage.file_path
    }
  } catch (e) {
    console.error('Ошибка загрузки группы:', e)
  }
})

function isValidPhone(phone) {
  return /^\d{11}$/.test(phone)
}

async function submitApplication() {
  const {
    child_surname,
    child_name,
    child_lastname,
    parent_surname,
    parent_name,
    parent_lastname,
    parent_phone
  } = form.value

  if (
    !child_surname || !child_name || !child_lastname ||
    !parent_surname || !parent_name || !parent_lastname || !parent_phone
  ) {
    errorMessage.value = 'Пожалуйста, заполните все поля.'
    return
  }

  if (!isValidPhone(parent_phone)) {
    errorMessage.value = 'Введите корректный номер телефона (11 цифр).'
    return
  }

  errorMessage.value = ''

  const payload = {
    group_id: route.params.id,
    child_surname,
    child_name,
    child_lastname,
    parent_surname,
    parent_name,
    parent_lastname,
    parent_phone
  }

  const res = await fetch('/api/register-child', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  })

  if (res.ok) {
    alert('Заявка отправлена!')
    resetForm()
    showForm.value = false
  } else {
    alert('Ошибка при отправке заявки.')
  }
}

function resetForm() {
  form.value = {
    child_surname: '',
    child_name: '',
    child_lastname: '',
    parent_surname: '',
    parent_name: '',
    parent_lastname: '',
    parent_phone: ''
  }
  errorMessage.value = ''
  showForm.value = false
}
</script>


<style scoped lang="sass">
.group-single
  padding: 40px 20px
  max-width: 960px
  margin: 0 auto

  &__content
    display: flex
    flex-direction: column
    padding-top: 20px
    gap: 24px

  &__image
    width: 100%
    max-height: 400px
    object-fit: cover
    border-radius: 12px

  &__title
    font-size: 32px
    font-weight: 700
    margin-bottom: 12px

  &__description
    font-size: 16px
    line-height: 1.6
    white-space: pre-line

.apply-button,
.cancel-button,
.apply-form button[type="submit"]
  margin-top: 30px
  padding: 0.75rem
  font-size: 1rem
  border: none
  border-radius: 6px
  cursor: pointer
  transition: background-color 0.3s

.apply-button,
.apply-form button[type="submit"]
  background-color: #1C5242
  color: white
  &:hover
    background-color: #1d7c60

.cancel-button
  background-color: #ccc
  color: #333
  margin-left: 8px
  &:hover
    background-color: #bbb

.apply-form
  display: flex
  flex-direction: column
  gap: 1rem

.apply-form__title
  font-size: 2rem
  margin-top: 40px
  margin-bottom: 10px
  color: #1C5242
  font-weight: bold

input
  padding: 0.75rem
  font-size: 1rem

.error-message
  color: red
  font-weight: bold

.teacher-card
  display: flex
  gap: 20px
  background: #f7f7f7
  padding: 20px
  border-radius: 12px
  margin-top: 20px
  align-items: center

  &__photo
    width: 120px
    height: 120px
    object-fit: cover
    border-radius: 12px

  &__info
    flex: 1

  &__name
    font-size: 20px
    font-weight: bold
    margin-bottom: 8px

  &__phone
    font-size: 16px
    color: #555
</style>
