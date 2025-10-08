# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| 1.1.x   | :white_check_mark: |
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of go-svg-captcha seriously. If you believe you have found a security vulnerability, please report it to us as described below.

### Please Do Not

- **Do not** open a public GitHub issue for security vulnerabilities
- **Do not** disclose the vulnerability publicly until it has been addressed

### Please Do

1. **Email** the maintainers directly at the email address listed in the GitHub profile
2. **Include** the following information:
   - Type of vulnerability (e.g., XSS, SVG injection, etc.)
   - Full paths of source file(s) related to the vulnerability
   - The location of the affected source code (tag/branch/commit or direct URL)
   - Any special configuration required to reproduce the issue
   - Step-by-step instructions to reproduce the issue
   - Proof-of-concept or exploit code (if possible)
   - Impact of the issue, including how an attacker might exploit it

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours
- **Communication**: We will keep you informed of our progress towards resolving the issue
- **Timeline**: We aim to address critical vulnerabilities within 7 days
- **Credit**: If you wish, we will credit you in the security advisory and release notes

## Security Best Practices for Users

When using go-svg-captcha in your applications, follow these security best practices:

### 1. Server-Side Validation

Always validate CAPTCHA responses on the server side:

```go
// Store CAPTCHA answer in secure server-side session
session.Values["captcha"] = result.Text
session.Values["captcha_created"] = time.Now()

// Validate user input
if userInput != session.Values["captcha"] {
    return errors.New("invalid CAPTCHA")
}
```

### 2. Single Use

Invalidate CAPTCHA after one attempt (success or failure):

```go
// After validation, immediately delete
delete(session.Values, "captcha")
session.Save(r, w)
```

### 3. Time Limits

Set expiration times for CAPTCHAs:

```go
createdAt := session.Values["captcha_created"].(time.Time)
if time.Since(createdAt) > 10*time.Minute {
    return errors.New("CAPTCHA expired")
}
```

### 4. Rate Limiting

Implement rate limiting to prevent brute force attacks:

```go
// Limit CAPTCHA generation per IP
// Limit validation attempts per session
```

### 5. HTTPS Only

Always serve CAPTCHAs over HTTPS to prevent man-in-the-middle attacks.

### 6. Secure Storage

Never store CAPTCHA answers in:

- Client-side cookies
- Hidden form fields
- Client-side JavaScript variables
- URL parameters

### 7. Content Security Policy

If embedding SVG directly in HTML, use appropriate CSP headers:

```
Content-Security-Policy: default-src 'self'; img-src 'self' data:;
```

### 8. Input Sanitization

While this library generates safe SVG output, always sanitize user input when validating:

```go
// Trim and validate input
userInput = strings.TrimSpace(userInput)
if len(userInput) > maxCaptchaLength {
    return errors.New("invalid input")
}
```

## Known Security Considerations

### SVG Security

The SVG output from this library is generated programmatically and does not include:

- External resources
- JavaScript code
- Event handlers
- External links

However, if you modify the SVG output, ensure you:

- Sanitize any external data
- Validate any injected content
- Use CSP headers appropriately

### CAPTCHA Strength

This library provides visual CAPTCHAs that are:

- **Suitable for** preventing basic automated bot attacks
- **Not suitable for** high-security applications requiring strong bot prevention
- **Not OCR-proof** against determined attackers with ML models

For high-security applications, consider:

- Using additional authentication factors
- Implementing more complex CAPTCHA challenges
- Using professional CAPTCHA services (reCAPTCHA, hCaptcha)

### Accessibility

CAPTCHAs can be a barrier for users with visual impairments. Consider:

- Providing alternative verification methods
- Implementing audio CAPTCHAs
- Using more accessible authentication approaches

## Security Updates

Security updates will be:

- Released as soon as possible after validation
- Announced in GitHub Security Advisories
- Included in CHANGELOG.md with `[SECURITY]` tag
- Communicated via GitHub releases

## Disclosure Policy

When we receive a security report:

1. **Confirmation** - We confirm the vulnerability and its impact
2. **Development** - We develop a patch
3. **Release** - We release the patched version
4. **Disclosure** - We publicly disclose the vulnerability after the patch is available

We follow responsible disclosure practices and will:

- Work with the reporter to understand and address the issue
- Keep the reporter informed throughout the process
- Credit the reporter (unless they prefer anonymity)

## Contact

For security concerns, contact:

- **GitHub Issues**: For non-security bugs only
- **Email**: See GitHub profile for contact information
- **Security Advisory**: Use GitHub's private vulnerability reporting feature

## Additional Resources

- [OWASP CAPTCHA Guide](https://cheatsheetseries.owasp.org/cheatsheets/CAPTCHA_Cheat_Sheet.html)
- [Go Security Best Practices](https://golang.org/doc/security/best-practices)

Thank you for helping keep go-svg-captcha and its users safe!
