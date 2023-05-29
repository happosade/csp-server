// Somewhat questionable if this is needed at all,
// or should this resource be actually protected by WAF.
// All public facing telemetry/logging tools are always
// prone to get malicious data and that's just something
// we have to take into account.
package middleware
