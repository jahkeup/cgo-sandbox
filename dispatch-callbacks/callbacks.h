#ifndef CALLBACKS_H_
#define CALLBACKS_H_

typedef void step_one_cb(void* data);
typedef step_one_cb* step_one_cb_func;

typedef void step_two_cb(void* data, const int x);
typedef step_two_cb* step_two_cb_func;

typedef void step_three_cb(void* data, const int x, const char*);
typedef step_three_cb* step_three_cb_func;

struct callbacks_unit {
    step_one_cb* step_one_cb;
    step_two_cb* step_two_cb;
    step_three_cb* step_three_cb;
};

typedef struct callbacks_unit callbacks_unit;

enum callbacks_result {
  CALLBACKS_SUCCESS,
  CALLBACKS_FAILURE,
};

int invoke_callbacks(callbacks_unit* cb, void* data);

#endif // CALLBACKS_H_
