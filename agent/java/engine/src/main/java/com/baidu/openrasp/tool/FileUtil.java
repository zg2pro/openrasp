/*
 * Copyright 2017-2021 Baidu Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.baidu.openrasp.tool;

import com.baidu.openrasp.EngineBoot;
import org.apache.commons.io.FileUtils;

import java.io.File;
import java.io.InputStream;
import java.io.OutputStream;
import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;

/**
 * Created by tyy on 4/7/17.
 * 文件工具类
 */
public class FileUtil {
    /**
     * 读取文件内容
     *
     * @param file 文件对象
     * @return 文件字符串内容
     * @throws IOException {@link IOException}
     */
    public static String readFileByFile(File file) throws IOException {
        return FileUtils.readFileToString(file);
    }

    /**
     * 还原文件真实路径,避免绕过
     *
     * @param file 文件
     * @return 真实文件路径su
     */
    public static String getRealPath(File file) {
        String absPath = file.getAbsolutePath();
        if (OSUtil.isWindows()) {
            int index = absPath.indexOf("::$");
            if (index >= 0) {
                file = new File(absPath.substring(0, index));
            }
        }
        try {
            return file.getCanonicalPath();
        } catch (IOException e) {
            return absPath;
        }
    }

    /**
     * 获取当前jar包所在的文件夹路径
     *
     * @return jar包所在文件夹路径
     */
    public static String getBaseDir() {
        String baseDir;
        String jarPath = EngineBoot.class.getProtectionDomain().getCodeSource().getLocation().getFile();
        try {
            baseDir = URLDecoder.decode(new File(jarPath).getParent(), "UTF-8");
        } catch (UnsupportedEncodingException e) {
            baseDir = new File(jarPath).getParent();
        }
        System.out.println("########### baseDir:"+baseDir);
        return baseDir;
    }

    public static void copyStream(InputStream in, OutputStream out, int bufferSize) throws Exception {
        byte[] buffer = new byte[bufferSize];
        int len = 0;

        while ((len = in.read(buffer)) > 0) {
            out.write(buffer, 0, len);
        }
    }

}
