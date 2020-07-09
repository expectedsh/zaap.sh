const { toString } = Object.prototype

const isFunction = (obj) => typeof (obj) === 'function'
const isObject = (obj) => obj === Object(obj)
const isArray = (obj) => toString.call(obj) === '[object Array]'
const isDate = (obj) => toString.call(obj) === '[object Date]'
const isRegExp = (obj) => toString.call(obj) === '[object RegExp]'
const isBoolean = (obj) => toString.call(obj) === '[object Boolean]'
const isNumerical = (obj) => !Number.isNaN(obj - 0)

function processKeys(convert, obj, path = []) {
  if (!isObject(obj) || isDate(obj) || isRegExp(obj) || isBoolean(obj) || isFunction(obj)) {
    return obj
  }

  if (isArray(obj)) {
    return obj.map((v, index) => processKeys(convert, v, [...path, index]))
  }

  return Object.fromEntries(
    Object.entries(obj)
      .filter(([key]) => Object.prototype.hasOwnProperty.call(obj, key))
      .map(([key, value]) => [convert(key, path), processKeys(convert, value, [...path, key])]),
  )
}

function camelize(options) {
  return (string, path) => {
    if (options?.exclude?.(path) || isNumerical(string)) {
      return string
    }
    string = string.replace(/[-_\s]+(.)?/g, (match, chr) => (chr ? chr.toUpperCase() : ''))
    // Ensure 1st char is always lowercase
    return string.substr(0, 1).toLowerCase() + string.substr(1)
  }
}

function decamelize(options) {
  return (string, path) => (options?.exclude?.(path)
    ? string
    : string.split(/(?=[A-Z])/).join('_').toLowerCase())
}

export function camelizeKeys(object, options) {
  return processKeys(camelize(options), object)
}

export function decamelizeKeys(object, options) {
  return processKeys(decamelize(options), object)
}
