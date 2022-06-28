/* eslint-disable @typescript-eslint/no-use-before-define */
import {
    b as e,
    w as t,
    d as o
}
from "./withConfigTransport.js";
import {
    ref as r,
    inject as a,
    onMounted as n,
    onUnmounted as l,
    computed as i,
    watch as s,
    resolveComponent as d,
    openBlock as p,
    createElementBlock as c,
    withKeys as v,
    withModifiers as u,
    Fragment as h,
    renderList as f,
    normalizeClass as m,
    normalizeStyle as g,
    createBlock as b,
    createCommentVNode as y,
    createElementVNode as w,
    toDisplayString as x,
    withDirectives as O,
    createVNode as k,
    vShow as C,
    defineComponent as S,
    nextTick as $,
    withCtx as _,
    createSlots as B,
    renderSlot as V,
    normalizeProps as I,
    guardReactiveProps as T
} from "vue";
import {
    flatten as A,
    isFunction as j,
    isString as M
}
from "lodash-es";
import {
    F
}
from "./consts.js";
import {
    g as H,
    h as N
}
from "./color-config.js";
import E, {
    VaDropdownContent as L
}
from "./index24.js";
import P from "./index27.js";
import D from "./index30.js";
import {
    __decorate as z
}
from "tslib";
import {
    V as U,
    m as K,
    O as R,
    p as q
}
from "./vue-class-component.esm-bundler.js";
import {
    C as W
}
from "./ColorMixin.js";
import {
    s as G
}
from "./style-inject.es.js";
import {
    V as J
}
from "./index55.js";
import "./VaConfig.js";
import "colortranslator";
import "asva-executors";
import "@popperjs/core";
import "./StatefulMixin.js";
import "./VaIcon.js";
import "./SizeMixin.js";
import "cleave.js";
const Q = {
    options: {
        type: Array,
        default: () => []
    },
    textBy: {
        type: [String, Function],
        default: "text"
    },
    valueBy: {
        type: [String, Function]
    },
    trackBy: {
        type: [String, Function],
        default: "value"
    },
    disabledBy: {
        type: [String, Function],
        default: "disabled"
    }
};
const X = {
        rules: {
            type: Array,
            default: () => []
        },
        disabled: {
            type: Boolean,
            default: !1
        },
        readonly: {
            type: Boolean,
            default: !1
        },
        success: {
            type: Boolean,
            default: !1
        },
        messages: {
            type: Array,
            default: []
        },
        error: {
            type: Boolean,
            default: !1
        },
        errorMessages: {
            type: [Array, String]
        },
        errorCount: {
            type: Number,
            default: 1
        },
        id: {
            type: [String, Number],
            default: void 0
        },
        name: {
            type: [String, Number],
            default: void 0
        },
        modelValue: {
            default: void 0,
            validator: () => {
                throw new Error("ValidateMixin: `modelValue` prop should be defined in component.")
            }
        }
    },
    prepareValidations = (e = [], t = null) => (M(e) && (e = [e]), e.map((e => j(e) ? e(t) : e)));
const Y = {
    loading: {
        type: Boolean,
        default: !1
    }
};
const Z = {
    maxSelections: {
        type: Number,
        default: void 0
    }
};
const ee = U.with(class SelectOptionListProps {
    constructor() {
        this.options = q({
            type: Array,
            default: () => []
        }), this.noOptionsText = q({
            type: String,
            default: "Items not found"
        }), this.getSelectedState = q({
            type: Function,
            default: () => !1
        }), this.getText = q({
            type: Function
        }), this.getTrackBy = q({
            type: Function
        }), this.multiple = q({
            type: Boolean,
            default: !1
        }), this.search = q({
            type: String,
            default: ""
        }), this.hoveredOption = q({
            type: [String, Object],
            default: null
        }), this.tabindex = q({
            type: Number,
            default: 0
        })
    }
});
let te = class VaSelectOptionList extends(K(W, ee)) {
    constructor() {
        // eslint-disable-next-line prefer-rest-params
        super(...arguments), this.itemRefs = {}
    }
    created() {
        s((() => this.$props.hoveredOption), (e => {
            e && this.scrollToOption(e)
        }))
    }
    beforeUpdate() {
        this.itemRefs = {}
    }
    setItemRef(e) {
        return t => {
            t && (this.itemRefs[e] = t)
        }
    }
    get hoveredOptionComputed() {
        return this.hoveredOption || null
    }
    set hoveredOptionComputed(e) {
        this.$emit("update:hoveredOption", e)
    }
    get filteredOptions() {
        return this.$props.search ? this.$props.options.filter((e => {
            const t = this.$props.getText(e).toString().toUpperCase(),
                o = this.$props.search.toUpperCase();
            return t.includes(o)
        })) : this.$props.options
    }
    selectOption(e) {
        this.$emit("select-option", e)
    }
    getOptionClass(e) {
        return {
            "va-select-option-list__option": !0,
            "va-select-option-list__option--selected": this.$props.getSelectedState(e)
        }
    }
    getOptionStyle(e) {
        return {
            color: this.$props.getSelectedState(e) ? this.colorComputed : "inherit",
            backgroundColor: this.isHovered(e) ? N(this.colorComputed) : "transparent"
        }
    }
    isHovered(e) {
        return !!this.hoveredOptionComputed && ("string" == typeof e ? e === this.hoveredOptionComputed : !!this.getTrackBy && this.getTrackBy(this.hoveredOptionComputed) === this.getTrackBy(e))
    }
    updateHoveredOption(e) {
        this.hoveredOptionComputed = e || null
    }
    hoverPreviousOption() {
        if (this.hoveredOptionComputed) {
            const e = this.filteredOptions.findIndex((e => this.$props.getText(e) === this.$props.getText(this.hoveredOptionComputed)));
            this.filteredOptions[e - 1] ? this.hoveredOptionComputed = this.filteredOptions[e - 1] : this.$emit("no-previous-option-to-hover")
        } else {
            this.filteredOptions.length && this.updateHoveredOption(this.filteredOptions[this.filteredOptions.length - 1])
        }
    }
    hoverNextOption() {
        if (this.hoveredOptionComputed) {
            const e = this.filteredOptions.findIndex((e => this.$props.getText(e) === this.$props.getText(this.hoveredOptionComputed)));
            this.filteredOptions[e + 1] && (this.hoveredOptionComputed = this.filteredOptions[e + 1])
        } else {
            this.filteredOptions.length && this.updateHoveredOption(this.filteredOptions[0])
        }
    }
    hoverFirstOption() {
        this.filteredOptions.length > 0 && this.updateHoveredOption(this.filteredOptions[0])
    }
    scrollToOption(e) {
        const t = this.itemRefs[this.$props.getTrackBy(e)];
        t && t.scrollIntoView({
            behavior: "auto",
            block: "nearest",
            inline: "nearest"
        })
    }
    focus() {
        this.$refs.el && this.$refs.el.focus()
    }
};
te = z([R({
    name: "VaSelectOptionList",
    components: {
        VaIcon: P
    },
    emits: ["select-option", "update:hoveredOption", "no-previous-option-to-hover"]
})], te);
const oe = te;
const re = ["tabindex"],
    ae = ["onClick", "onMouseover"],
    ne = {
        key: 1,
        class: "va-select-option-list no-options"
    };
G(':root{--va-gray-light:#acb5be;--va-light-gray:#eee;--va-light-gray2:#eff4f5;--va-light-gray3:#f5f8f9;--va-lighter-gray:#ddd;--va-charcoal:#555;--va-darkest-gray:#333;--va-almost-black:#161616;--va-hover-black:#222;--va-vue-green:#4ae387;--va-vue-light-green:#dbf9e8;--va-light-green:#c8f9c5;--va-lighter-green:#d6ffd3;--va-light-blue:#dcf1ff;--va-light-yellow:#fff1c8;--va-light-pink:#ffcece;--va-vue-darkest-blue:#34495e;--va-vue-turquoise:#dbf9e7;--va-white:#fff;--va-theme-danger:#e34b4a;--va-theme-warning:#ffc200;--va-theme-red:#e34a4a;--va-theme-orange:#f7cc36;--va-theme-blue:#4ab2e3;--va-theme-blue-dark:#2c82e0;--va-theme-violet:#db76df;--va-theme-pale:#d9d9d9;--va-primary:var(--va-vue-green);--va-danger:var(--va-theme-red);--va-warning:var(--va-theme-orange);--va-info:var(--va-theme-blue);--va-success:var(--va-vue-green);--va-secondary:#babfc2;--va-dark-gray:#282828;--va-gray:#adb3b9;--va-default-gray:#8396a5;--va-dark-blue:#0e4ac4;--va-text-gray:#b4b4b4;--va-markdown-code:#d7234e;--va-prism-background:#f4f8fa;--va-separator-color:#e6e9ec;--va-border:0;--va-font-family:"Source Sans Pro",sans-serif;--va-letter-spacing:0.0375rem;--va-background-color:#fff;--va-swing-transition:0.3s cubic-bezier(0.25,0.8,0.5,1);--va-block-border-radius:0.375rem;--va-block-border:thin solid rgba(52,56,85,0.25);--va-block-box-shadow:0 2px 3px 0 rgba(52,56,85,0.25);--va-control-box-shadow:none;--va-control-border:0;--va-transition:0.2s cubic-bezier(0.4,0,0.6,1);--va-outline-border-width:0.125rem;--va-outline-box-shadow:none;--va-square-border-radius:0.25rem;--va-form-padding:1.25rem;--va-form-border-radius:0.125rem;--va-text-selected:#b3d4fc;--va-text-highlighted:#fff3d1;--va-link-color:var(--va-primary);--va-link-color-secondary:var(--va-secondary);--va-link-color-hover:var(--va-primary);--va-link-color-active:var(--va-primary);--va-link-color-visited:var(--va-primary);--va-muted:#7f828b;--va-li-background:var(--va-theme-blue-dark);--va-text-block:var(--va-light-gray3);--va-stripe-border-size:0.25rem;--va-box-shadow:0 0.25rem 0.5rem 0 rgba(59,63,73,0.15);--va-select-option-list-display:flex;--va-select-option-list-flex-direction:column;--va-select-option-list-width:100%;--va-select-option-list-list-style:none;--va-select-option-list-option-cursor:pointer;--va-select-option-list-option-display:flex;--va-select-option-list-option-align-items:center;--va-select-option-list-option-padding:0.375rem 12px;--va-select-option-list-option-min-height:2.25rem;--va-select-option-list-option-word-break:break-word;--va-select-option-list-icon-margin-right:0.5rem;--va-select-option-list-selected-icon-margin-left:auto;--va-select-option-list-selected-icon-font-size:0.8rem}.va-select-option-list{display:var(--va-select-option-list-display);flex-direction:var(--va-select-option-list-flex-direction);list-style:var(--va-select-option-list-list-style);max-height:200px;width:var(--va-select-option-list-width)}.va-select-option-list__option{align-items:var(--va-select-option-list-option-align-items);cursor:var(--va-select-option-list-option-cursor);display:var(--va-select-option-list-option-display);min-height:var(--va-select-option-list-option-min-height);padding:var(--va-select-option-list-option-padding);word-break:var(--va-select-option-list-option-word-break)}.va-select-option-list__option--icon{margin-right:var(--va-select-option-list-icon-margin-right)}.va-select-option-list__option--selected-icon{font-size:var(--va-select-option-list-selected-icon-font-size);margin-left:var(--va-select-option-list-selected-icon-margin-left)}.va-select-option-list.no-options{padding:.5rem}'), oe.render = function render$1(e, t, o, r, a, n) {
    const l = d("va-icon");
    return p(), c("div", {
        class: "va-select-option-list",
        ref: "el",
        tabindex: e.tabindex,
        onKeydown: [t[0] || (t[0] = v(u(((...t) => e.hoverPreviousOption && e.hoverPreviousOption(...t)), ["stop", "prevent"]), ["up"])), t[1] || (t[1] = v(u(((...t) => e.hoverPreviousOption && e.hoverPreviousOption(...t)), ["stop", "prevent"]), ["left"])), t[2] || (t[2] = v(u(((...t) => e.hoverNextOption && e.hoverNextOption(...t)), ["stop", "prevent"]), ["down"])), t[3] || (t[3] = v(u(((...t) => e.hoverNextOption && e.hoverNextOption(...t)), ["stop", "prevent"]), ["right"]))]
    }, [e.filteredOptions.length ? (p(!0), c(h, {
        key: 0
    }, f(e.filteredOptions, (t => (p(), c("div", {
        key: e.$props.getTrackBy(t),
        ref: e.setItemRef(e.$props.getTrackBy(t)),
        class: m(e.getOptionClass(t)),
        style: g(e.getOptionStyle(t)),
        onClick: u((o => e.selectOption(t)), ["stop"]),
        onMouseover: o => e.updateHoveredOption(t)
    }, [t.icon ? (p(), b(l, {
        key: 0,
        size: "small",
        class: "va-select-option-list__option--icon",
        name: t.icon
    }, null, 8, ["name"])) : y("v-if", !0), w("span", null, x(e.getText(t)), 1), O(k(l, {
        class: "va-select-option-list__option--selected-icon",
        size: "small",
        name: "done",
        color: e.colorComputed
    }, null, 8, ["color"]), [
        [C, e.$props.getSelectedState(t)]
    ])], 46, ae)))), 128)) : (p(), c("div", ne, x(e.noOptionsText), 1))], 40, re)
}, oe.__file = "src/components/va-select/VaSelectOptionList/VaSelectOptionList.vue";
const le = S({
    name: "VaSelect",
    components: {
        VaSelectOptionList: t(oe),
        VaIcon: P,
        VaDropdown: E,
        VaDropdownContent: L,
        VaInput: D,
        VaInputWrapper: J
    },
    emits: ["update-search", "update:modelValue", "clear"],
    props: {...Q,
        ...X,
        ...Y,
        ...Z,
        modelValue: {
            type: [String, Number, Object, Array],
            default: ""
        },
        position: {
            type: String,
            default: "bottom",
            validator: e => ["top", "bottom"].includes(e)
        },
        allowCreate: {
            type: [Boolean, String],
            default: !1,
            validator: e => [!0, !1, "unique"].includes(e)
        },
        color: {
            type: String,
            default: "primary"
        },
        multiple: {
            type: Boolean,
            default: !1
        },
        searchable: {
            type: Boolean,
            default: !1
        },
        disabled: {
            type: Boolean,
            default: !1
        },
        readonly: {
            type: Boolean,
            default: !1
        },
        separator: {
            type: String,
            default: ", "
        },
        width: {
            type: String,
            default: "100%"
        },
        maxHeight: {
            type: String,
            default: "128px"
        },
        clearValue: {
            type: String,
            default: ""
        },
        noOptionsText: {
            type: String,
            default: "Items not found"
        },
        fixed: {
            type: Boolean,
            default: !0
        },
        clearable: {
            type: Boolean,
            default: !1
        },
        clearableIcon: {
            type: String,
            default: "highlight_off"
        },
        hideSelected: {
            type: Boolean,
            default: !1
        },
        tabindex: {
            type: Number,
            default: 0
        },
        dropdownIcon: {
            type: [String, Object],
            default: () => ({
                open: "expand_more",
                close: "expand_less"
            }),
            validator: e => {
                if ("string" == typeof e) {
                    return !0
                }
                const t = "string" == typeof e.open,
                    o = "string" == typeof e.close;
                return t && o
            }
        },
        outline: {
            type: Boolean,
            default: !1
        },
        bordered: {
            type: Boolean,
            default: !1
        },
        label: {
            type: String,
            default: ""
        },
        placeholder: {
            type: String,
            default: ""
        }
    },
    setup(t, d) {
        const p = r(),
            c = r(),
            v = r(),
            {
                getOptionByValue: u,
                getValue: h,
                getText: f,
                getTrackBy: m
            } = function useSelectableList(t) {
                const getValue = o => t.valueBy ? "string" == typeof o ? o : e(o, t.valueBy) : o;
                return {
                    isSelectableListComponent: !0,
                    getValue: getValue,
                    getOptionByValue: e => t.valueBy && t.options.find((t => e === getValue(t))) || e,
                    getText: o => "string" == typeof o || "number" == typeof o ? o : e(o, t.textBy),
                    getDisabled: o => "string" != typeof o && e(o, t.disabledBy),
                    getTrackBy: o => "string" == typeof o ? o : e(o, t.trackBy)
                }
            }(t),
            {
                validate: g,
                isFocused: b,
                computedErrorMessages: y,
                computedError: w
            } = function useFormComponent(e, t) {
                const o = r(!1),
                    s = r(!1),
                    d = r([]),
                    p = r(!1),
                    c = r(!0),
                    v = a(F, void 0);
                n((() => {
                    (null == v ? void 0 : v.onChildMounted) && v.onChildMounted(t)
                })), l((() => {
                    (null == v ? void 0 : v.onChildUnmounted) && v.onChildUnmounted(t)
                }));
                // eslint-disable-next-line @typescript-eslint/no-use-before-define
                const validate = () => (h.value = !1, f.value = [], e.rules && e.rules.length > 0 && prepareValidations(A(e.rules), e.modelValue).forEach((e => {
                        // eslint-disable-next-line @typescript-eslint/no-use-before-define
                        M(e) ? (f.value.push(e), h.value = !0) : !1 === e && (h.value = !0)
                            // eslint-disable-next-line @typescript-eslint/no-use-before-define
                    })), !h.value),
                    u = i((() => o.value)),
                    h = i({
                        get: () => e.error || p.value,
                        set(e) {
                            p.value = e
                        }
                    }),
                    f = i({
                        get: () => e.errorMessages ? prepareValidations(e.errorMessages) : d.value,
                        set(e) {
                            d.value = e
                        }
                    });
                return {
                    isFocused: s,
                    isFormComponent: c,
                    formProvider: v,
                    validate: validate,

                    // eslint-disable-next-line @typescript-eslint/camelcase
                    ValidateMixin_onBlur: () => {
                        s.value = !1,
                            h.value = !1,
                            validate()
                    },
                    shouldValidateOnBlur: u,
                    focus: () => {
                        throw new Error("focus method should be implemented in the component")
                    },
                    reset: () => {
                        throw new Error("reset method should be implemented in the component")
                    },
                    resetValidation: () => {
                        f.value = [],
                            h.value = !1
                    },
                    hasError: () => h.value,
                    computedError: h,
                    computedErrorMessages: f
                }
            }(t, d),
            {
                colorComputed: x
            } = function useColor(e) {
                const t = {
                        getColor: H
                    },
                    o = i((() => t.getColor(e.color)));
                return {
                    hasColorTheme: !0,
                    theme: t,
                    colorComputed: o,
                    computeColor: (e, o) => t.getColor(e, o)
                }
            }(t),
            O = r(""),
            k = i((() => t.searchable || t.allowCreate));
        s((() => O.value), (e => {
            d.emit("update-search", e),
                j.value = null
        }));
        const C = i({
                get() {
                    const e = u(t.modelValue);
                    return t.multiple ? e ? Array.isArray(e) ? e : [e] : [] : Array.isArray(e) && (o("Model value should be a string for a single Select."), e.length) ? e[e.length - 1] : e
                },
                set(e) {
                    d.emit("update:modelValue", h(e))
                }
            }),
            S = i((() => C.value ? "string" == typeof C.value ? C.value : Array.isArray(C.value) ? C.value.map((e => f(e))).join(t.separator) || t.clearValue : f(C.value) : t.clearValue)),
            _ = i((() => !!t.clearable && (!t.disabled && (t.multiple ? !!C.value.length : C.value !== t.clearValue)))),
            // eslint-disable-next-line @typescript-eslint/no-use-before-define
            B = i((() => t.dropdownIcon ? "string" == typeof t.dropdownIcon ? t.dropdownIcon : N.value ? t.dropdownIcon.close : t.dropdownIcon.open : "")),
            // eslint-disable-next-line @typescript-eslint/no-use-before-define
            V = i((() => t.options ? t.hideSelected ? t.options.filter((e => !checkIsOptionSelected(e))) : t.options : [])),
            // eslint-disable-next-line @typescript-eslint/no-use-before-define
            checkIsOptionSelected = e => !!C.value && (Array.isArray(C.value) ? !!C.value.find((t => compareOptions(t, e))) : compareOptions(C.value, e)),
            compareOptions = (e, t) => e === t || ("string" == typeof e && "string" == typeof t ? e === t : null !== e && null !== t && ("object" == typeof e && "object" == typeof t && m(e) === m(t))),
            {
                exceedsMaxSelections: I,
                addOption: T
            } = function useMaxSelections(e, t, o) {
                return {
                    exceedsMaxSelections: () => void 0 !== t.value && e.value.length >= t.value,
                    addOption: t => {
                        const r = [...e.value, t];
                        o("update:modelValue", r)
                    }
                }
            }(C, r(t.maxSelections), d.emit),
            selectOption = e => {
                // eslint-disable-next-line @typescript-eslint/no-use-before-define
                if (null !== j.value) {
                    if (k.value && (O.value = ""), t.multiple) {
                        if (checkIsOptionSelected(e)) { C.value = C.value.filter((t => !compareOptions(e, t))) } else {
                            if (I()) { return }
                            T(e)
                        }
                    } else {
                        // eslint-disable-next-line @typescript-eslint/no-use-before-define
                        C.value = "string" == typeof e ? e : {...e }, hideAndFocus()
                    }
                    // eslint-disable-next-line @typescript-eslint/no-use-before-define
                } else { hideAndFocus() }
            },
            allowedToCreate = () => !(!t.allowCreate || "" === O.value),
            addNewOption = () => {
                if (t.multiple) {
                    if (I()) { return }
                    const e = C.value.some((e => e === O.value));
                    "unique" === t.allowCreate && e || (C.value = [...C.value, O.value])
                } else { C.value = O.value }
                O.value = ""
            },
            j = r(null),
            selectHoveredOption = () => {
                // eslint-disable-next-line @typescript-eslint/no-use-before-define
                N.value ? selectOption(j.value) : showDropdown()
            },
            N = r(!1),
            E = i({
                get: () => N.value,
                set: e => {
                    // eslint-disable-next-line @typescript-eslint/no-use-before-define
                    e ? showDropdown() : hideDropdown()
                }
            }),
            L = i((() => !(t.multiple || t.searchable || t.allowCreate))),
            showDropdown = () => {
                N.value = !0,
                    // eslint-disable-next-line @typescript-eslint/no-use-before-define
                    scrollToSelected(),
                    focusSearchOrOptions()
            },
            hideDropdown = () => {
                N.value = !1,
                    g()
            },
            toggleDropdown = () => {
                N.value ? hideAndFocus() : showDropdown()
            },
            P = i((() => b.value || N.value)),
            hideAndFocus = () => {
                let e;
                hideDropdown(),
                    null === (e = p.value) || void 0 === e || e.focus()
            },
            focusSearchBar = () => {
                let e;
                null === (e = v.value) || void 0 === e || e.focus()
            },
            focusOptionList = () => {
                let e;
                null === (e = c.value) || void 0 === e || e.focus()
            },
            focusSearchOrOptions = () => {
                $((() => {
                    k.value ? focusSearchBar() : focusOptionList()
                }))
            },
            D = i((() => t.disabled ? -1 : t.tabindex)),
            scrollToSelected = () => {
                const e = C.value;
                if (!e.length && "object" != typeof e) { return }
                const t = Array.isArray(e) ? e[e.length - 1] : e;
                j.value = t,
                    $((() => {
                        let e;
                        return null === (e = c.value) || void 0 === e ? void 0 : e.scrollToOption(t)
                    }))
            };
        let z, U = "";
        const K = ["ArrowUp", "ArrowDown", "ArrowLeft", "ArrowRight", "Enter", " "];
        return {
            select: p,
            optionList: c,
            focusOptionList: focusOptionList,
            focus: () => {
                t.disabled || (b.value = !0)
            },
            blur: () => {
                b.value = !1,
                    g()
            },
            reset: () => {
                t.multiple ? C.value = Array.isArray(t.clearValue) ? t.clearValue : [] : C.value = t.clearValue,
                    O.value = "",
                    d.emit("clear")
            },
            onSelectClick: () => {
                t.disabled || toggleDropdown()
            },
            hideAndFocus: hideAndFocus,
            searchBar: v,
            focusSearchBar: focusSearchBar,
            searchInput: O,
            showSearchInput: k,
            hoveredOption: j,
            tabIndexComputed: D,
            valueComputed: C,
            valueComputedString: S,
            showClearIcon: _,
            toggleIcon: B,
            showDropdownContent: N,
            computedErrorMessages: y,
            computedError: w,
            filteredOptions: V,
            checkIsOptionSelected: checkIsOptionSelected,
            closeOnContentClick: L,
            selectOption: selectOption,
            selectOrAddOption: () => {
                null === j.value ? allowedToCreate() && addNewOption() : selectHoveredOption()
            },
            selectHoveredOption: selectHoveredOption,
            hoverPreviousOption: () => {
                let e;
                null === (e = c.value) || void 0 === e || e.hoverPreviousOption()
            },
            hoverNextOption: () => {
                let e;
                null === (e = c.value) || void 0 === e || e.hoverNextOption()
            },
            showDropdownContentComputed: E,
            showDropdown: showDropdown,
            hideDropdown: hideDropdown,
            toggleDropdown: toggleDropdown,
            isFocusedComputed: P,
            colorComputed: x,
            onHintedSearch: e => {
                if (K.some((t => t === e.key))) { return }
                const o = 1 === e.key.length,
                    r = "Backspace" === e.key || "Delete" === e.key;
                if (clearTimeout(z), r ? U = U ? U.slice(0, -1) : "" : o && (U += e.key), k.value) { O.value = U } else {
                    if (U) {
                        const e = t.options.find((e => f(e).toLowerCase().startsWith(U.toLowerCase())));
                        e && (j.value = e)
                    }
                    z = setTimeout((() => {
                        U = ""
                    }), 1e3)
                }
            },
            getText: f,
            getTrackBy: m
        }
    }
});
const ie = ["tabindex"],
    se = {
        class: "va-input__append"
    },
    de = {
        class: "va-select-dropdown__options-wrapper"
    };
G(':root{--va-gray-light:#acb5be;--va-light-gray:#eee;--va-light-gray2:#eff4f5;--va-light-gray3:#f5f8f9;--va-lighter-gray:#ddd;--va-charcoal:#555;--va-darkest-gray:#333;--va-almost-black:#161616;--va-hover-black:#222;--va-vue-green:#4ae387;--va-vue-light-green:#dbf9e8;--va-light-green:#c8f9c5;--va-lighter-green:#d6ffd3;--va-light-blue:#dcf1ff;--va-light-yellow:#fff1c8;--va-light-pink:#ffcece;--va-vue-darkest-blue:#34495e;--va-vue-turquoise:#dbf9e7;--va-white:#fff;--va-theme-danger:#e34b4a;--va-theme-warning:#ffc200;--va-theme-red:#e34a4a;--va-theme-orange:#f7cc36;--va-theme-blue:#4ab2e3;--va-theme-blue-dark:#2c82e0;--va-theme-violet:#db76df;--va-theme-pale:#d9d9d9;--va-primary:var(--va-vue-green);--va-danger:var(--va-theme-red);--va-warning:var(--va-theme-orange);--va-info:var(--va-theme-blue);--va-success:var(--va-vue-green);--va-secondary:#babfc2;--va-dark-gray:#282828;--va-gray:#adb3b9;--va-default-gray:#8396a5;--va-dark-blue:#0e4ac4;--va-text-gray:#b4b4b4;--va-markdown-code:#d7234e;--va-prism-background:#f4f8fa;--va-separator-color:#e6e9ec;--va-border:0;--va-font-family:"Source Sans Pro",sans-serif;--va-letter-spacing:0.0375rem;--va-background-color:#fff;--va-swing-transition:0.3s cubic-bezier(0.25,0.8,0.5,1);--va-block-border-radius:0.375rem;--va-block-border:thin solid rgba(52,56,85,0.25);--va-block-box-shadow:0 2px 3px 0 rgba(52,56,85,0.25);--va-control-box-shadow:none;--va-control-border:0;--va-transition:0.2s cubic-bezier(0.4,0,0.6,1);--va-outline-border-width:0.125rem;--va-outline-box-shadow:none;--va-square-border-radius:0.25rem;--va-form-padding:1.25rem;--va-form-border-radius:0.125rem;--va-text-selected:#b3d4fc;--va-text-highlighted:#fff3d1;--va-link-color:var(--va-primary);--va-link-color-secondary:var(--va-secondary);--va-link-color-hover:var(--va-primary);--va-link-color-active:var(--va-primary);--va-link-color-visited:var(--va-primary);--va-muted:#7f828b;--va-li-background:var(--va-theme-blue-dark);--va-text-block:var(--va-light-gray3);--va-stripe-border-size:0.25rem;--va-box-shadow:0 0.25rem 0.5rem 0 rgba(59,63,73,0.15);--va-select-cursor:pointer;--va-select-dropdown-border-radius:4px;--va-select-dropdown-background:#fff;--va-select-scroll-color:grey;--va-select-box-shadow:0 4px 8px rgba(59,63,73,0.15)}.va-select,.va-select .va-input{cursor:var(--va-select-cursor)}.va-select .va-input__append{align-content:center;display:flex;justify-content:center}.va-select-dropdown .va-dropdown__anchor{display:block}.va-select-dropdown__content{border-bottom-left-radius:var(--va-select-dropdown-border-radius);border-bottom-right-radius:var(--va-select-dropdown-border-radius);border-top-left-radius:0;border-top-right-radius:0;box-shadow:var(--va-select-box-shadow);overflow:hidden;padding:0}.va-select-dropdown__options-wrapper{background:var(--va-select-dropdown-background);overflow-y:auto;scrollbar-color:var(--va-select-scroll-color) transparent;scrollbar-width:thin}.va-select-dropdown__options-wrapper::-webkit-scrollbar{width:4px}.va-select-dropdown__options-wrapper::-webkit-scrollbar-track{border-radius:10px;box-shadow:none}.va-select-dropdown__options-wrapper::-webkit-scrollbar-thumb{background:var(--va-select-scroll-color);border-radius:2px;opacity:.3}'), le.render = function render(e, t, o, r, a, n) {
    const l = d("va-icon"),
        i = d("va-input"),
        s = d("va-select-option-list"),
        c = d("va-dropdown-content"),
        h = d("va-dropdown"),
        f = d("va-input-wrapper");
    return p(), b(f, {
        success: e.$props.success,
        messages: e.$props.messages,
        error: e.$props.error,
        "error-messages": e.computedErrorMessages,
        style: g({
            width: e.$props.width
        })
    }, {
        default: _((() => [k(h, {
            ref: "dropdown",
            modelValue: e.showDropdownContentComputed,
            "onUpdate:modelValue": t[17] || (t[17] = t => e.showDropdownContentComputed = t),
            position: e.$props.position,
            disabled: e.$props.disabled,
            "max-height": e.$props.maxHeight,
            fixed: e.$props.fixed,
            "close-on-content-click": e.closeOnContentClick,
            trigger: "none",
            class: "va-select__dropdown va-select-dropdown",
            "keep-anchor-width": "",
            "boundary-body": "",
            stateful: !1
        }, {
            anchor: _((() => [w("div", {
                class: "va-select",
                ref: "select",
                tabindex: e.tabIndexComputed,
                onFocus: t[0] || (t[0] = (...t) => e.focus && e.focus(...t)),
                onBlur: t[1] || (t[1] = (...t) => e.blur && e.blur(...t)),
                onKeydown: [t[2] || (t[2] = v(u((t => e.onSelectClick()), ["stop", "prevent"]), ["enter"])), t[3] || (t[3] = v(u((t => e.onSelectClick()), ["stop", "prevent"]), ["space"]))],
                onClick: t[4] || (t[4] = u((t => e.onSelectClick()), ["prevent"]))
            }, [y(" We show messages outside of dropdown to draw dropdown content under the input "), k(i, {
                "model-value": e.valueComputedString,
                success: e.$props.success,
                error: e.computedError,
                clearable: e.showClearIcon,
                clearableIcon: e.$props.clearableIcon,
                color: e.$props.color,
                label: e.$props.label,
                placeholder: e.$props.placeholder,
                loading: e.$props.loading,
                disabled: e.$props.disabled,
                outline: e.$props.outline,
                bordered: e.$props.bordered,
                focused: e.isFocusedComputed,
                tabindex: -1,
                readonly: "",
                onCleared: e.reset
            }, B({
                appendInner: _((() => [w("div", se, [e.$slots.appendInner ? V(e.$slots, "appendInner", {
                    key: 0
                }) : y("v-if", !0), k(l, {
                    color: e.colorComputed,
                    name: e.toggleIcon
                }, null, 8, ["color", "name"])])])),
                _: 2
            }, [e.$slots.prepend ? {
                name: "prepend",
                fn: _((() => [V(e.$slots, "prepend")]))
            } : void 0, e.$slots.append ? {
                name: "append",
                fn: _((() => [V(e.$slots, "append")]))
            } : void 0, e.$slots.prependInner ? {
                name: "prependInner",
                fn: _((() => [V(e.$slots, "prependInner")]))
            } : void 0, e.$slots.content ? {
                name: "content",
                fn: _((({
                    value: t,
                    focus: o
                }) => [V(e.$slots, "content", I(T({
                    valueString: t,
                    focus: o,
                    value: e.valueComputed
                })))]))
            } : void 0]), 1032, ["model-value", "success", "error", "clearable", "clearableIcon", "color", "label", "placeholder", "loading", "disabled", "outline", "bordered", "focused", "onCleared"])], 40, ie)])),
            default: _((() => [k(c, {
                class: "va-select-dropdown__content",
                // eslint-disable-next-line @typescript-eslint/no-empty-function
                onKeyup: t[16] || (t[16] = v(u((() => {}), ["stop"]), ["enter"])),
                onKeydown: [v(u(e.hideAndFocus, ["prevent"]), ["esc"]), v(e.hideDropdown, ["tab"])]
            }, {
                default: _((() => [e.showSearchInput ? (p(), b(i, {
                    key: 0,
                    id: e.$props.id,
                    ref: "searchBar",
                    modelValue: e.searchInput,
                    "onUpdate:modelValue": t[5] || (t[5] = t => e.searchInput = t),
                    class: "va-select__input",
                    placeholder: "Search",
                    removable: "",
                    name: e.$props.name,
                    tabindex: e.tabindex + 1,
                    bordered: !0,
                    onKeydown: [t[6] || (t[6] = v(u((t => e.hoverPreviousOption()), ["stop", "prevent"]), ["up"])), t[7] || (t[7] = v(u((t => e.hoverPreviousOption()), ["stop", "prevent"]), ["left"])), t[8] || (t[8] = v(u((t => e.hoverNextOption()), ["stop", "prevent"]), ["down"])), t[9] || (t[9] = v(u((t => e.hoverNextOption()), ["stop", "prevent"]), ["right"])), t[10] || (t[10] = v(u((t => e.selectOrAddOption()), ["prevent"]), ["enter"]))],
                    onFocus: t[11] || (t[11] = t => e.hoveredOption = null)
                }, null, 8, ["id", "modelValue", "name", "tabindex"])) : y("v-if", !0), w("div", de, [k(s, {
                    ref: "optionList",
                    hoveredOption: e.hoveredOption,
                    "onUpdate:hoveredOption": t[12] || (t[12] = t => e.hoveredOption = t),
                    style: g({
                        maxHeight: e.$props.maxHeight
                    }),
                    options: e.filteredOptions,
                    "selected-value": e.valueComputed,
                    "get-selected-state": e.checkIsOptionSelected,
                    "get-text": e.getText,
                    "get-track-by": e.getTrackBy,
                    search: e.searchInput,
                    "no-options-text": e.$props.noOptionsText,
                    color: e.$props.color,
                    tabindex: e.tabindex + 1,
                    onSelectOption: e.selectOption,
                    onNoPreviousOptionToHover: t[13] || (t[13] = t => e.focusSearchBar()),
                    onKeydown: [t[14] || (t[14] = v(u((t => e.selectHoveredOption()), ["stop", "prevent"]), ["enter"])), t[15] || (t[15] = v(u((t => e.selectHoveredOption()), ["stop", "prevent"]), ["space"])), e.onHintedSearch]
                }, null, 8, ["hoveredOption", "style", "options", "selected-value", "get-selected-state", "get-text", "get-track-by", "search", "no-options-text", "color", "tabindex", "onSelectOption", "onKeydown"])])])),
                _: 1
            }, 8, ["onKeydown"])])),
            _: 3
        }, 8, ["modelValue", "position", "disabled", "max-height", "fixed", "close-on-content-click"])])),
        _: 3
    }, 8, ["success", "messages", "error", "error-messages", "style"])
}, le.__file = "src/components/va-select/VaSelect.vue";
const pe = t(oe);
const ce = t(le);
export {
    pe as VaSelectOptionList, ce as
    default
};
//# sourceMappingURL=index41.js.map