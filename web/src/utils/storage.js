/* eslint-disable no-constant-condition */

/**
 * storage handler
 */

const SessionUserKey = 'LOGIN_USER';
const SessionAdminKey = 'LOGIN_ADMIN';

export function getSessionId() {
    const user = getLoginUser();
    return user ? user['session_id'] : '';
}

export function getLoginAdmin() {
    const value = sessionStorage.getItem(SessionAdminKey);
    if (value) {
        return JSON.parse(value);
    } else {
        return null;
    }
}

export function getLoginUser() {
    const value = sessionStorage.getItem(SessionUserKey);
    if (value) {
        return JSON.parse(value);
    } else {
        return null;
    }
}

export function setLoginUser(user) {
    sessionStorage.setItem(SessionUserKey, JSON.stringify(user))
}

export function setLoginAdmin(admin) {
    sessionStorage.setItem(SessionAdminKey, JSON.stringify(admin))
}
