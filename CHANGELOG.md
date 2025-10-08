# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned

- Generate multiple CAPTCHAs in batch
- Export to PNG/JPEG format
- Audio CAPTCHA for accessibility
- More math operations (multiplication, division)

## [1.1.0] - 2024-03-26

### Added

- Noise generation for enhanced security and visual complexity
- Custom font support via `FontPath` option
- Users can now use their own TTF fonts instead of the default embedded font

### Fixed

- Random color generation not working correctly
- Color algorithm now properly generates diverse colors with appropriate contrast
- HSL to RGB conversion improved for better color distribution

### Changed

- Improved color generation algorithm for better visual appearance
- Enhanced character distortion for increased security

## [1.0.5] - 2023-12-22

### Added

- Complete package documentation with examples
- GoDoc comments for all exported types and functions
- Usage examples in README

### Changed

- Removed custom text option from public API for consistency
- Removed size option from math CAPTCHA (not applicable)
- Simplified API surface for better user experience

### Fixed

- API inconsistencies between text and math CAPTCHA options

## [1.0.4] - 2023-12-21

### Added

- Initial public release of go-svg-captcha
- Text-based CAPTCHA generation with customizable options
- Math-based CAPTCHA generation (addition and subtraction)
- SVG output without external image processing dependencies
- Configurable width, height, and font size
- Grayscale and color mode support
- Inverse mode (light text on dark background)
- Background color customization
- Curve distortion for security
- Custom character set support
- Character exclusion for readability (e.g., remove O, 0, I, 1)
- Embedded default font (Comismsh)

---

## Release Links

- [1.1.0](https://github.com/reu98/go-svg-captcha/releases/tag/v1.1.0) - 2024-03-26
- [1.0.5](https://github.com/reu98/go-svg-captcha/releases/tag/v1.0.5) - 2023-12-22
- [1.0.4](https://github.com/reu98/go-svg-captcha/releases/tag/v1.0.4) - 2023-12-21

---

## Categories Legend

- **Added** - New features
- **Changed** - Changes in existing functionality
- **Deprecated** - Soon-to-be removed features
- **Removed** - Removed features
- **Fixed** - Bug fixes
- **Security** - Security improvements
