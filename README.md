## This project has been discontinued

Yes, it still works, you can still use this software. It just won't recieve any updates now.

> I do not have the interest or time to oversee the development of this software. I thank all the amazing people who contributed to this project and made what it is in it's final form.

# Roop

> Take a video and replace the face in it with a face of your choice. You only need one image of the desired face. No dataset, no training.

[![Build Status](https://img.shields.io/github/actions/workflow/status/s0md3v/roop/ci.yml.svg?branch=main)](https://github.com/s0md3v/roop/actions?query=workflow:ci)

<img src="https://i.ibb.co/4RdPYwQ/Untitled.jpg"/>

## Installation

Be aware, the installation needs technical skills and is not for beginners. Please do not open platform and installation related issues on GitHub.

[Basic](https://github.com/s0md3v/roop/wiki/1.-Installation) - It is more likely to work on your computer, but will be quite slow

[Acceleration](https://github.com/s0md3v/roop/wiki/2.-Acceleration) - Unleash the full potential of your CPU and GPU


## Usage

Start the program with arguments:

```
python run.py [options]

-h, --help                                                                 show this help message and exit 显示帮助信息后退出
-s SOURCE_PATH, --source SOURCE_PATH                                       select an source image 最终期望的人脸图片
-t TARGET_PATH, --target TARGET_PATH                                       select an target image or video 需要换成的图片或视频
-o OUTPUT_PATH, --output OUTPUT_PATH                                       select output file or directory 最终文件生成目录
--frame-processor FRAME_PROCESSOR [FRAME_PROCESSOR ...]                    frame processors (choices: face_swapper, face_enhancer, ...) 帧处理器,默认：face_swapper (可多选: face_swapper, face_enhancer)
--keep-fps                                                                 keep target fps 保持原视频帧率
--keep-frames                                                              keep temporary frames 保持原视频帧文件目录
--skip-audio                                                               skip target audio 跳过视频音频
--many-faces                                                               process every face 处理多张脸
--reference-face-position REFERENCE_FACE_POSITION                          position of the reference face 人脸位置参考：用于定位人脸
--reference-frame-number REFERENCE_FRAME_NUMBER                            number of the reference frame 人脸帧数参考：用于定位人脸
--similar-face-distance SIMILAR_FACE_DISTANCE                              face distance used for recognition 人脸识别阈值，默认值：0.85。对比人脸的相识度，遇到跳脸现象，可加大阈值
--temp-frame-format {jpg,png}                                              image format used for frame extraction 帧文件图片格式
--temp-frame-quality [0-100]                                               image quality used for frame extraction 帧文件质量
--output-video-encoder {libx264,libx265,libvpx-vp9,h264_nvenc,hevc_nvenc}  encoder used for the output video 输出视频编码
--output-video-quality [0-100]                                             quality used for the output video 输出视频质量
--max-memory MAX_MEMORY                                                    maximum amount of RAM in GB 最大内存(以GB为单位)
--execution-provider {cpu} [{cpu} ...]                                     available execution provider (choices: cpu, ...) 可用的执行提供程序,默认:cpu ,通常使用cpu和cuda (需要安装相关环境)
--execution-threads EXECUTION_THREADS                                      number of execution threads 执行线程数
-v, --version                                                              show program's version number and exit 显示程序的版本号并退出
```


### Headless

Using the `-s/--source`, `-t/--target` and `-o/--output` argument will run the program in headless mode.


## Disclaimer

This software is designed to contribute positively to the AI-generated media industry, assisting artists with tasks like character animation and models for clothing.

We are aware of the potential ethical issues and have implemented measures to prevent the software from being used for inappropriate content, such as nudity.

Users are expected to follow local laws and use the software responsibly. If using real faces, get consent and clearly label deepfakes when sharing. The developers aren't liable for user actions.


## Licenses

Our software uses a lot of third party libraries as well pre-trained models. The users should keep in mind that these third party components have their own license and terms, therefore our license is not being applied.


## Credits

- [deepinsight](https://github.com/deepinsight) for their [insightface](https://github.com/deepinsight/insightface) project which provided a well-made library and models.
- all developers behind the libraries used in this project


## Documentation

Read the [documentation](https://github.com/s0md3v/roop/wiki) for a deep dive.

# 额外需要安装

+ 下载[visualstudio](https://visualstudio.microsoft.com/zh-hans/visual-cpp-build-tools/)
需要完整安装c++部分
+ ffmpeg
+ 下载[模型文件](https://cdn-lfs.huggingface.co/repos/f7/d7/f7d7d940dac099fff8d8e8a39c964e80ea9a852a5d0f4926ef9d25d2dda69273/e4a3f08c753cb72d04e10aa0f7dbe3deebbf39567d4ead6dce08e98aa49e16af?response-content-disposition=inline%3B+filename*%3DUTF-8%27%27inswapper_128.onnx%3B+filename%3D%22inswapper_128.onnx%22%3B&Expires=1726417750&Policy=eyJTdGF0ZW1lbnQiOlt7IkNvbmRpdGlvbiI6eyJEYXRlTGVzc1RoYW4iOnsiQVdTOkVwb2NoVGltZSI6MTcyNjQxNzc1MH19LCJSZXNvdXJjZSI6Imh0dHBzOi8vY2RuLWxmcy5odWdnaW5nZmFjZS5jby9yZXBvcy9mNy9kNy9mN2Q3ZDk0MGRhYzA5OWZmZjhkOGU4YTM5Yzk2NGU4MGVhOWE4NTJhNWQwZjQ5MjZlZjlkMjVkMmRkYTY5MjczL2U0YTNmMDhjNzUzY2I3MmQwNGUxMGFhMGY3ZGJlM2RlZWJiZjM5NTY3ZDRlYWQ2ZGNlMDhlOThhYTQ5ZTE2YWY%7EcmVzcG9uc2UtY29udGVudC1kaXNwb3NpdGlvbj0qIn1dfQ__&Signature=c-5fVZSuIEQp5skCQFmTDZ3HhdFbNJzuf6Qx2c-W1GLgekYxA1AtpEv1Wk2r%7EBThB2p1p12yWHQbpQLV0IDbKA2jTha%7EZdqxm0-ZqvuYmtCqzmN8C3ZILN7udUPs%7EeTb6mTfECoqwVDhZEsKvhycdUsHomdHZxFgaFB%7EVNlbLkpcT9-Y%7EqhA4C%7EH8TMVcpHX1qi2qAq43-PBdWK%7ErlZLXIeqAgXdzwRTchQ3Uifn1dFOYfIyjGYFhDUFwLyqyIQqW7qmwZo2rVQDq6Y9cgx%7EVbnNPcqAvXgdVM0CcqmKng4tnt9TpbPbkuzZknoGcQaUQxK3Q2Udceox%7EIPY6oF-FA__&Key-Pair-Id=K3ESJI6DHPFC7)到`..\models\inswapper_128.onnx`

# 启动命令

`python run.py --execution-provider cuda`

# 进入虚拟环境

`.\.venv\Scripts\activate.ps1`

# 退出虚拟环境

`deactivate`