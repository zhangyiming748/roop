要将 Python 项目打包为可在其他计算机上运行的 `.exe` 文件，您可以使用 `PyInstaller` 或 `cx_Freeze` 等工具。以下是使用 `PyInstaller` 的步骤，因为它是最常用的工具之一。

### 步骤 1: 创建虚拟环境

如果您还没有创建虚拟环境，可以使用以下命令创建并激活它：

```bash
# 创建虚拟环境
python -m venv venv

# 激活虚拟环境
# Windows
venv\Scripts\activate

# macOS/Linux
source venv/bin/activate
```

### 步骤 2: 安装 PyInstaller

在激活的虚拟环境中，安装 `PyInstaller`：

```bash
pip install pyinstaller
```

### 步骤 3: 打包项目

使用 `PyInstaller` 打包您的 Python 脚本。假设您的主脚本名为 `your_script.py`，您可以运行以下命令：

```bash
pyinstaller --onefile your_script.py
```

- `--onefile` 选项会将所有依赖项打包到一个单独的可执行文件中。
- 您还可以添加其他选项，例如 `--windowed`（对于 GUI 应用程序，避免显示命令行窗口）。

### 步骤 4: 查找生成的可执行文件

打包完成后，您将在项目目录中看到一个 `dist` 文件夹，里面包含生成的 `.exe` 文件。您可以在 `dist` 文件夹中找到 `your_script.exe`。

### 步骤 5: 测试可执行文件

在您的计算机上测试生成的 `.exe` 文件，确保它按预期工作。

### 步骤 6: 分发可执行文件

将 `dist` 文件夹中的 `.exe` 文件复制到其他计算机上。通常，您只需将 `.exe` 文件复制到目标计算机上即可运行，而不需要安装 Python 或其他依赖项。

### 注意事项

- 确保在打包时，您的代码没有依赖于特定的文件路径或环境变量。
- 如果您的应用程序使用了外部文件（如配置文件、数据文件等），您可能需要将这些文件一起分发，并在代码中处理它们的路径。
- 在不同的操作系统上打包时，确保在相应的操作系统上运行 `PyInstaller`，因为生成的 `.exe` 文件只能在 Windows 上运行。

如果您在打包过程中遇到任何问题，请提供具体的错误信息或问题描述，以便我能更好地帮助您！