/**
 * Define all of your application routes here
 * for more information on routes,
 * official documentation https://router.vuejs.org/en/
 */
import DashboardView from '@/views/DashboardView.vue'
import LoginHome from '@/views/LoginHome.vue'
import LoginForm from '@/views/LoginForm.vue'
import Userweb from '@/views/UserWeb.vue'
import Menu from '@/views/Menu.vue'
import Role from '@/views/Role.vue'
import UserRole from '@/views/UserRole'
import RoleMenu from '@/views/RoleMenu'
import Home from '@/views/Home.vue'
import Hello from '@/views/Hello.vue'
import Page404 from '@/views/404.vue'
import Page403 from '@/views/403.vue'
import UserLogin from '@/views/UserLogin.vue'
import SandBox from '@/views/SandBox.vue'

import { chatUrl, frontendUrl, clientId, loggingUrl, ssoUrl, userManagementUrl } from '../config/env'

export default [
  {
    path: '/',
    component: LoginHome,
    children: [
      { path: '', redirect: '/home-page' },
      {
        path: '/home-page',
        name: 'HOME',
        component: LoginForm,
        meta: { loginForm: true },
      },
      {
        path: '/logout',
        name: 'HOME',
        component: LoginForm,
        meta: { loginForm: true },
      },
      {
        path: '/oauth/code',
        name: 'LOGIN_HOME',
        component: LoginForm,
        props: route => ({ code: route.query.code }),
        meta: { loginForm: true },
      },
      {
        path: '/forgot-password',
        name: 'FORGOT_PASSWORD',
        beforeEnter (to, from, next) {
          if (from.path === '/') {
            alert('You will be redirected to Forgot Password Menu in User Management System')
            window.location = `${userManagementUrl}/forgot-password`
          } else {
            window.open(`${userManagementUrl}/forgot-password`, '_blank')
          }
        },
      },
      {
        path: '/404',
        name: '404',
        component: Page404,
        meta: { Page404: true },
      },
      {
        path: '/403',
        name: '403',
        component: Page403,
        meta: { Page403: true },
      },
    ],
  },
  {
    path: '/dashboard-view',
    component: DashboardView,
    children: [
      {
        path: '/home',
        name: 'HOME',
        component: Home,
        meta: { requiresAuth: true },
      },
      {
        path: '/hello',
        name: 'Hello',
        component: Hello,
        meta: { requiresAuth: true },
      },
      {
        path: '/sandbox',
        name: 'SANDBOX',
        component: SandBox,
        meta: { requiresAuth: false },
      },
      {
        path: '/userweb',
        name: 'USERWEB',
        component: Userweb,
        meta: { requiresAuth: true, kode: 'USER' },
      },
      {
        path: '/menu',
        name: 'MENU',
        component: Menu,
        meta: { requiresAuth: true },
      },
      {
        path: '/role',
        name: 'ROLE',
        component: Role,
        meta: { requiresAuth: true },
      },
      {
        path: '/user-role',
        name: 'USER_ROLE',
        component: UserRole,
        meta: { requiresAuth: true },
      },
      {
        path: '/role-menu',
        name: 'ROLE_MENU',
        component: RoleMenu,
        meta: { requiresAuth: true },
      },
      {
        path: '/change-password',
        name: 'CHANGE_PASSWORD',
        beforeEnter (to, from, next) {
          if (from.path === '/') {
            alert('You will be redirected to Change Password Menu in User Management System')
            window.location = `${userManagementUrl}/change-password`
          } else {
            window.open(`${userManagementUrl}/change-password`, '_blank')
          }
        },
      },

      {
        path: '/profil-pengguna',
        name: 'PROFIL_PENGGUNA',
        beforeEnter (to, from, next) {
          if (from.path === '/') {
            alert('You will be redirected to User Profile Menu in User Management System')
            window.location = `${userManagementUrl}/profil-pengguna`
          } else {
            window.open(`${userManagementUrl}/profil-pengguna`, '_blank')
          }
        },
      },

      {
        path: '/pengguna-login',
        name: 'PENGGUNA_LOGIN',
        component: UserLogin,
        meta: { requiresAuth: true },
      },
    ],
  },
]
