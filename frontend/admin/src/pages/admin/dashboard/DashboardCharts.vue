<template>

  <div class="row row-equal">
    <div class="flex xs9">
      <va-card color="divider">
        <va-card-title>
          {{ $t('titles.carouselPreview') }}
        </va-card-title>
        <va-card-content>
          <el-carousel
            height="32rem"
            indicator-position="outside"
          >
            <el-carousel-item v-for="(image, index) in images" :key="index" >
              <div class="image">
                <el-image
                  fit="cover"
                  :src="image.src"
                >
                </el-image>
              </div>
            </el-carousel-item>
          </el-carousel>
        </va-card-content>
      </va-card>
    </div>
    <div class="flex xs3 manage">
      <div class="manage-list">
        <va-card color="divider">
          <va-card-title>
            {{ $t('titles.carouselManage') }}
          </va-card-title>
          <va-card-content>
            <div class="table-wrapper">
              <table class="va-table">
                <thead>
                  <tr>
                    <th>{{ $t('tables.headings.title') }}</th>
                    <th>{{ $t('tables.headings.operation') }}</th>
                  </tr>
                </thead>
                <tbody>
                  <!-- eslint-disable-next-line vue/no-unused-vars -->
                  <tr v-for="(image, index) in images" :key="index">
                    <td>
                      <p> {{ image.title }} </p>
                    </td>
                    <td>
                      <va-icon-up class="icon" @click="upBanner(index)" />
                      <va-icon-down class="icon" @click="downBanner(index)" />
                      <va-icon-edit class="icon" @click="modifyDialog=true, selectedUUID=image.uuid, selectedTitle=image.title, selectedIndex=index" />
                      <va-icon-delete class="icon" @click="deleteBanner(index)" />
                    </td>
                  </tr>
                </tbody>
              </table> 
            </div>
          </va-card-content>
        </va-card>
      </div>

      <div class="manage-button">
        <va-card color="divider">
          <va-card-content>

            <va-chip class="mb-2 mr-2"
              small="true"
              color="#87CEEB"
              @click="addDialog='true'"
            >
              <el-upload
                action="/api/image/upload"
                :show-file-list="false"
                :on-success="handleSuccess"
              >
                <span class="white-color">
                  {{ $t('forms.inputs.upload') }}
                </span>
              </el-upload>
            </va-chip>

            <va-chip class="mb-2 mr-2 buttons"
              small="true"
              color="#87CEEB"
              @click="reset()"
            >
              <span class="white-color">
                {{ $t('buttons.reset')}}
              </span>
            </va-chip>

            <va-chip class="mb-2 mr-2 buttons"
              small="true"
              color="#87CEEB"
              @click="apply()"
            >
              <span class="white-color">
                {{ $t('buttons.apply')}}
              </span>
            </va-chip>
          </va-card-content>
        </va-card>
      </div>
    </div>
  </div>

  <va-modal
    v-model="modifyDialog"
    position="top"
    size="large"
    :title=" $t('modal.tips')"
    :message=" $t('modal.modifyTitleMessage')"
    :okText=" $t('modal.confirm')"
    :cancelText=" $t('modal.cancel')"
    @ok="setTitle()"
  >
    <va-input
      v-model="selectedTitle"
      @keypress.enter="setTitle()"
      :placeholder='$t("modal.title")'
      class="va-chat__input mr-2"
    />
  </va-modal>

</template>

<script>
import VaIconDelete from '@/components/icons/VaIconDelete.vue'
import VaIconUp from '@/components/icons/VaIconUp.vue'
import VaIconDown from '@/components/icons/VaIconDown.vue'
import VaIconEdit from '@/components/icons/VaIconEdit.vue'
import VaChart from '@/components/va-charts/VaChart.vue'
import DashboardContributorsChart from './DashboardContributorsList'
import { useGlobalConfig } from 'vuestic-ui'
import { reactive, onMounted } from 'vue'
import { get, post } from '@/axios/html'

export default {
  name: 'dashboard-charts',
  // eslint-disable-next-line vue/no-unused-components
  components: { DashboardContributorsChart, VaChart, VaIconDelete, VaIconUp, VaIconDown, VaIconEdit },
  setup() {
    const state = reactive({
      images: [],
      backupImages: [],
      imagesQueue: [],
    })
    onMounted(async() => {
      get("/api/image/getImageUUID", '')
        .then(res => {
          if(res.data.code === '200') {
            res.data.UUIDs.forEach(UUID => {
              const src = "/api/image/getImage?uuid=" + UUID.uuid
              state.images.push({
                "subscript": UUID.subscript,
                "title": UUID.title,
                "uuid": UUID.uuid,
                "src": src
              })
              state.backupImages.push({
                "subscript": UUID.subscript,
                "title": UUID.title,
                "uuid": UUID.uuid,
                "src": src
              })
              state.imagesQueue.push({
                "subscript": UUID.subscript,
                "uuid": UUID.uuid,
              })
            })
          state.images.sort((a, b) => {
            return a.subscript - b.subscript
          })
          state.backupImages.sort((a, b) => {
            return a.subscript - b.subscript
          })
          state.imagesQueue.sort((a, b) => {
            return a.subscript - b.subscript
          })
          }
        }).catch(err => {
          console.log(err)
        })
    })
    return state
  },
  data () {
    return {
      modifyDialog: false,
      selectedUUID: '',
      selectedTitle: '',
      selectedIndex: -1,
      lineChartData: null,
      donutChartData: null,
      lineChartFirstMonthIndex: 0,
      toastColor: '#87CEEB',
      toastText: 'nullText',
      toastPosition: 'top-right',
      toastDuration: 2500,
    }
  },
  methods: {

    launchToast() {
      this.$vaToast.init({
        message: this.toastText,
        position: this.toastPosition,
        color: this.toastColor,
        duration: Number(this.toastDuration)
      })
    },
    switchBanner(firstIndex, secIndex) {
      const tempBanner = this.images[firstIndex]
      this.images[firstIndex] = this.images[secIndex]
      this.images[secIndex] = tempBanner

      const tempSubscript = this.imagesQueue[firstIndex]
      this.imagesQueue[firstIndex] = this.imagesQueue[secIndex]
      this.imagesQueue[secIndex] = tempSubscript

    },
    upBanner(index) {
      if(index !== 0) {
        this.images[index].subscript = index - 1
        this.switchBanner(index, index - 1)
      }
    },
    downBanner(index) {
      if(index !== (this.images.length - 1)) {
        this.images[index].subscript = index + 1
        this.switchBanner(index, index + 1)
      }
    },
    deleteBanner(index) {
      this.imagesQueue[index].subscript = -1
      this.images.splice(index, 1)
    },
    setTitle() {
      get("/api/image/setImageTitle", {
        uuid: this.selectedUUID,
        title: this.selectedTitle
      }).then(res => {
        if(res.data.code === '200') {
          this.images[this.selectedIndex].title = this.selectedTitle
          this.toastText = res.data.msg
          this.launchToast()
        }
      }).catch(err => {
        console.log(err)
      })
    },
    apply() {
      const indexArray = []
      this.imagesQueue.forEach(image => {
        if(image.subscript === -1) {
          indexArray.push({
            "index": -1,
            "uuid": image.uuid,
          })
        }else {
          indexArray.push({
            "index": this.imagesQueue.indexOf(image),
            "uuid": image.uuid,
          })
        }
      })
      post("/api/image/updateImageIndex", { 
        images: JSON.stringify(indexArray)
        })
        .then(res => {
          console.log(res)
          if(res.code === '200') {
            this.toastText = res.msg
            this.launchToast()
          }
        }).catch(err => {
          console.log(err)
        })
    },
    reset() {
      this.images = []
      this.backupImages.forEach(image => {
        this.images.push(image)
      })
    },
    handleSuccess(res) {
      if(res.code === '200') {
        const src = "/api/image/getImage?uuid=" + res.uuid
        this.images.push({
          "subscript": res.subscript,
          "title": res.title,
          "uuid": res.uuid,
          "src": src
        })
        this.backupImages.push({
          "subscript": res.subscript,
          "title": res.title,
          "uuid": res.uuid,
          "src": src
        })
        this.imagesQueue.push({
          "subscript": res.subscript,
          "uuid": res.uuid,
        })
        this.toastText = res.msg
        this.launchToast()
      } 
    }
  },
  computed: {
    theme() {
      return useGlobalConfig().getGlobalConfig().colors
    },
    donutChartDataURL () {
      return document.querySelector('.chart--donut canvas').toDataURL('image/png')
    },
  },
}
</script>

<style scoped>
  .table-wrapper {
    width: 100%;
    height: 38vh;
    margin: auto;
  }
  .table-wrapper th {
    position: sticky;
    top: 0;
    left: 0;
    background: #e1e9f8;
  }
  .table-wrapper table {
    word-wrap: break-word; 
    word-break: break-all;
    text-align: center;
    table-layout: fixed;
    width: 100%;
    margin: 0 auto;
    border-spacing: unset;
    border-top: unset;
  }
  .chart {
    height: 400px;
  }
  .text-right {
    text-align: right;
  }
  .operation {
    float: right;
  }
  .icon {
    margin-right: 5px;
  }
  .image {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
  }
  .buttons {
    float: right;
  }
  .white-color {
    color: white;
  }
  .manage-list {
    height: 32.8rem;
  }
  .manage-button {
    margin-top: 1.5rem;
    height: 4rem;
  }
</style>
