#include "callbacks.h"

int invoke_callbacks(callbacks_unit* cb, void* data) {
    if (cb->step_one_cb) {
        cb->step_one_cb(data);
    }
    if (cb->step_two_cb) {
        cb->step_two_cb(data, 2);
    }
    if (cb->step_three_cb) {
        cb->step_three_cb(data, 3, "three");
    }

    return CALLBACKS_SUCCESS;
}
