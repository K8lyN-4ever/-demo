<template>
  <div class="markup-tables flex">
    <va-card color="divider" class="flex mb-4">
      <va-card-content>
        <div class="table-wrapper">
          <va-chip class="mb-2 mr-2 add-button"
            small="true"
            color="#87CEEB"
            @click="addDialog='true'"
          >
            <span class="white-color">
              {{ $t('buttons.add')}}
            </span>
          </va-chip>
          <table class="va-table">
            <thead>
              <tr>
                <th>{{ $t('tables.headings.account') }}</th>
                <th>{{ $t('tables.headings.password') }}</th>
                <th>{{ $t('tables.headings.auth') }}</th>
                <th>{{ $t('tables.headings.operation') }}</th>
              </tr>
            </thead>
            <tbody>
              <!-- eslint-disable-next-line vue/no-unused-vars -->
              <tr v-for="(user, index) in users" :key="user.id">
                <td>
                  <auto-ellipsis :text="user.account" />
                </td>
                <td>
                  <auto-ellipsis :text="user.password" />
                </td>
                <td>
                  <el-select size="small" collapse-tags v-model='user.permission' placeholder='请选择' @change='setPermission(user)'>
                    <el-option v-for='(item, index) in auths' :key='index' :label='item' :value='item'></el-option>
                  </el-select>
                </td>
                <td >
                  <va-icon-edit class="icon" :title="$t('titles.edit')" v-if="user != ''" @click="editDialog=true, selectedUser=user, selectedIndex=index" />
                  <va-icon-delete class="icon" :title="$t('titles.delete')" v-if="user != ''" @click="deleteDialog=true, selectedUser=user, selectedIndex=index" />
                </td>
              </tr>
            </tbody>
          </table>          
        </div>
      </va-card-content>
    </va-card>
  </div>
  
  <va-modal
    v-model="deleteDialog"
    position="top"
    size="large"
    :title=" $t('modal.tips')"
    :message=" $t('modal.deleteAdminMessage')"
    :okText=" $t('modal.delete')"
    :cancelText=" $t('modal.cancel')"
    @ok="deleteRow"
  />

  <va-modal
    v-model="addDialog"
    position="top"
    size="large"
    :title=" $t('modal.addTips')"
    :message=" $t('modal.addMessage')"
    :okText=" $t('modal.add')"
    :cancelText=" $t('modal.cancel')"
    @ok="addUser"
  >
    <va-input
      v-model="account"
      @keypress.enter="addUser"
      :placeholder='$t("modal.account")'
      class="va-chat__input mr-2"
    />
    <br>
    <va-input
      v-model="password"
      @keypress.enter="addUser"
      :placeholder='$t("modal.password")'
      class="va-chat__input mr-2"
    />
    <br>
  </va-modal>

  <va-modal
    v-model="editDialog"
    position="top"
    size="large"
    :title=" $t('modal.editTips')"
    :message=" $t('modal.addMessage')"
    :okText=" $t('modal.confirm')"
    :cancelText=" $t('modal.cancel')"
    @ok="updatePassword"
  >
    <va-input
      v-model="account"
      @keypress.enter="updatePassword"
      :placeholder="this.selectedUser.account"
      readonly
      class="va-chat__input mr-2"
    />
    <br>
    <va-input
      v-model="password"
      @keypress.enter="updatePassword"
      :placeholder='$t("modal.password")'
      class="va-chat__input mr-2"
    />
    <br>
  </va-modal>
</template>

<script>
import VaIconDelete from '@/components/icons/VaIconDelete.vue'
import VaIconEdit from '@/components/icons/VaIconEdit.vue'
import AutoEllipsis from '@/pages/admin/ui/txt/AutoEllipsis.vue'
import { reactive, onMounted } from 'vue'
import { get } from '@/axios/html'

export default {
  components: {
    VaIconDelete,
    VaIconEdit,
    AutoEllipsis,
  },
  setup() {
    const state = reactive({
      users: [],
    })
    onMounted(async() => {
      get("/api/admin/getUsers", '')
        .then(res => {
          if(res.data.code === "0") {
            res.data.data.riders.forEach(option => {
              let auth = ''
              if(option.Type === 1) {
                auth = '管理员'
              }else if(option.Type === 0) {
                auth = '骑手'
              }
              state.users.push({
                "account": option.Account,
                "password": option.Password,
                'permission': auth
              })
           });
          }
        })
        .catch(err => {
          console.log(err)
        })
    })
    return state
  },
  data () {
    return {
       addDialog: false,
       editDialog: false,
       deleteDialog: false,
       account: "",
       password: "",
       selectedUser: [],
       selectedIndex: -1,
       auths: ["管理员", "骑手"],
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
    deleteRow() {
      get("/api/admin/deleteUser", {
        account: this.selectedUser.account
      })
        .then(res => {
          if(res.data.code === '0') {
            this.users.splice(this.selectedIndex, 1)
          }else {
            console.log("删除失败")
          }
          this.toastText = res.data.msg
          this.launchToast()
        })
        .catch(err => {
          console.log(err)
        })
    },
    addUser() {
      get('/api/user/register', {
        account: this.account,
        password: this.password
      }).then(res => {
        if(res.data.code === '0') {
          this.users.push({
            account: this.account,
            password: this.password,
            auth: '骑手'
          })
        }
        this.toastText = res.data.msg
        this.launchToast()
        this.account = ''
        this.password = ''
      }).catch(err => {
        console.log(err)
      })
    },
    setPermission(user) {
      let auth
      if(user.permission === '管理员') {
        auth = 1
      }else if(user.permission === '骑手') {
        auth = 0
      }
      get('/api/admin/updateUser', {
        account: user.account,
        type: auth
      }).then(res => {
        this.toastText = res.data.msg
        this.launchToast()
      }).catch(err => {
        console.log(err)
      })
    },
    updatePassword() {
      get('/api/admin/updatePassword',{
        account: this.selectedUser.account,
        password: this.password
      }).then(res => {
        if(res.data.code === "0") {
          this.selectedUser.password = this.password
        }
        this.password = ""
        this.toastText = res.data.msg
        this.launchToast()
      }).catch(err => {
        console.log(err)
      })
    }
  },
}
</script>

<style lang="scss">


  .markup-tables {

    *::-webkit-scrollbar {
      position: fixed;
      top: 50px;
      /*滚动条整体样式*/
      width: 5px;
      height: 2px;
    }

    *::-webkit-scrollbar-thumb {
      /*滚动条里面小方块*/
      background-color: rgba(0, 0, 0, 0.2);
      border-radius: 5px;
    }

    *::-webkit-scrollbar-track {
      /*滚动条里面轨道*/
      background-color: unset;
      border-radius: 5px;
    }

    .table-wrapper {
      width: 100%;
      height: 81vh;
      margin: auto;
      overflow-y: auto;
      background: #e1e9f8;
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
      // height: 200px;
      margin: 0 auto;
      border-spacing: unset;
      border-top: unset;
    }

    .table-wrapper table#table > thead > tr:nth-of-type(odd) {
      background: rgba(204, 204, 204, 0.4);
    }

    .va-table {
      // overflow: auto;
      width: 100%;
      // background: #D3D3D3;
    }

    td {
      min-width: 500px;
    }

    tr {
      border-collapse: separate;
      border-radius: 50px;
    }

    tr:hover {
      background: #F0F8FF;
    }

    .white-color {
      color: white;
    }

    .add-button {
      float: right;
    }

    .icon {
      margin-right: 10px;
    }

    .select {
      height: 22px;
      color: black;
    }
  }
</style>
