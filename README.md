## Changelog Usage

The changelog is a file that contains a curated, chronologically ordered list of notable changes for each version of a project. It helps users and contributors understand what has changed between versions.

### How to Use the Changelog

1. **Locate the Changelog**: The changelog file is typically named `CHANGELOG.md` and is located in the root directory of the project.

2. **Read the Changelog**:
    - Open the `CHANGELOG.md` file to view the list of changes.
    - The changelog is organized by version numbers, with the most recent changes at the top.
    - Each version entry includes a list of changes categorized by type (e.g., Added, Changed, Deprecated, Removed, Fixed, Security).

3. **Understand the Changes**:
    - **Added**: New features or functionalities.
    - **Changed**: Updates to existing features.
    - **Deprecated**: Features that are still available but will be removed in future releases.
    - **Removed**: Features that have been removed.
    - **Fixed**: Bug fixes.
    - **Security**: Security-related changes.

4. **Use the Information**:
    - Use the changelog to determine if you need to update your project to a new version.
    - Check for any breaking changes or deprecated features that might affect your project.
    - Review the new features and improvements to take advantage of them in your project.

### Example Changelog Entry

# core

git-chglog --next-tag v0.15.2 > CHANGELOG.md