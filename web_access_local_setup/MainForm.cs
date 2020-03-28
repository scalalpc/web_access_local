// Copyright 2020 The web_access_local Authors. All rights reserved.
// Use of this source code is governed by a GPL-style license that can be found in the LICENSE file.

using Microsoft.Win32;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.IO;
using System.Reflection;
using System.Text;
using System.Threading;
using System.Windows.Forms;

namespace web_access_local_setup
{
    public partial class MainForm : Form
    {
        private int _port;
        private string _authUrl;
        private string _dirPath;

        public MainForm()
        {
            InitializeComponent();

            this.pnlSetup.Parent = this;
            this.lblAuthUrl.Parent = this.pnlSetup;
            this.lblDirPath.Parent = this.pnlSetup;
            this.lblLocalPort.Parent = this.pnlSetup;
            this.txtAuthUrl.Parent = this.pnlSetup;
            this.txtDirPath.Parent = this.pnlSetup;
            this.txtLocalPort.Parent = this.pnlSetup;
            this.btnSelectDirPath.Parent = this.pnlSetup;
            this.btnSetup.Parent = this.pnlSetup;

            this.pnlFinish.Parent = this;
            this.pnlFinish.Visible = false;
            this.lblFinished.Parent = this.pnlFinish;
            this.btnClose.Parent = this.pnlFinish;
        }

        private void MainForm_Load(object sender, EventArgs e)
        {
        }

        private void btnSelectDirPath_Click(object sender, EventArgs e)
        {
            var result = this.dlgDirPath.ShowDialog();
            if (result == DialogResult.OK)
            {
                this.txtDirPath.Text = this.dlgDirPath.SelectedPath;
            }
        }

        private bool CheckInput()
        {
            if (!int.TryParse(this.txtLocalPort.Text, out _port) || _port <= 0 || _port > ushort.MaxValue)
            {
                MessageBox.Show("local port is invalid.");
                this.txtLocalPort.Focus();
                return false;
            }

            _authUrl = this.txtAuthUrl.Text.Trim();

            _dirPath = this.txtDirPath.Text.Trim();
            if (_dirPath.Length == 0)
            {
                MessageBox.Show("setup directory is invalid.");
                this.btnSelectDirPath.Focus();
                return false;
            }

            return true;
        }

        private void btnSetup_Click(object sender, EventArgs e)
        {
            bool valid = CheckInput();
            if (valid)
            {
                ChangeStatus("ready");

                var thread = new Thread(RunSetup);
                thread.IsBackground = true;
                thread.Start();
            }
        }

        private void RunSetup()
        {
            try
            {
                if (!Directory.Exists(_dirPath))
                {
                    Directory.CreateDirectory(_dirPath);
                }

                var embedFileName = "web_access_local.exe";
                var appName = embedFileName.Replace(".exe", "");
                var embedFilePath = Path.Combine(_dirPath, embedFileName);

                using (var stream = Assembly.GetExecutingAssembly().GetManifestResourceStream(string.Format("{0}.{1}", AppDomain.CurrentDomain.FriendlyName.Replace(".exe", ""), embedFileName)))
                {
                    if (File.Exists(embedFilePath))
                    {
                        File.Delete(embedFilePath);
                    }

                    var buffer = new byte[1024];
                    var size = 0;
                    using (var fs = File.OpenWrite(embedFilePath))
                    {
                        while (true)
                        {
                            size = stream.Read(buffer, 0, buffer.Length);
                            if (size > 0)
                            {
                                fs.Write(buffer, 0, size);
                            }
                            else
                            {
                                break;
                            }
                        }
                        fs.Close();
                    }

                    stream.Close();
                }

                var iniFileName = embedFileName + ".ini";

                var iniBuilder = new StringBuilder();
                iniBuilder.AppendLine("[app]");
                iniBuilder.Append("port=");
                iniBuilder.AppendLine(_port.ToString());
                iniBuilder.Append("authUrl=");
                iniBuilder.AppendLine(_authUrl);
                var iniContent = iniBuilder.ToString();

                var iniFilePath = Path.Combine(_dirPath, iniFileName);
                if (File.Exists(iniFilePath))
                {
                    File.Delete(iniFilePath);
                }
                using (var fs = File.OpenWrite(iniFilePath))
                {
                    var byteArr = Encoding.UTF8.GetBytes(iniContent);
                    fs.Write(byteArr, 0, byteArr.Length);
                    fs.Close();
                }

                using (var lmRegKey = Registry.LocalMachine)
                {
                    using (var runRgeKey = lmRegKey.CreateSubKey(@"Software\Microsoft\Windows\CurrentVersion\Run"))
                    {
                        runRgeKey.SetValue(appName, embedFilePath);
                        runRgeKey.Close();
                    }
                    lmRegKey.Close();
                }

                ChangeStatus("finish");
                ChangePanelVisibleForFinished();
            }
            catch (Exception ex)
            {
                ChangeStatus(ex.Message);
            }
        }

        private void btnClose_Click(object sender, EventArgs e)
        {
            Application.Exit();
        }

        private delegate void ChangeStatusDelegate(string text);

        private void ChangeStatus(string text)
        {
            if (this.lblStatus.InvokeRequired)
            {
                this.lblStatus.Invoke(new ChangeStatusDelegate(ChangeStatus), new object[] { text });
            }
            else
            {
                this.lblStatus.Text = text;
            }
        }

        private delegate void ChangePanelVisibleForFinishedDelegate();

        private void ChangePanelVisibleForFinished()
        {
            if (this.pnlSetup.InvokeRequired)
            {
                this.pnlSetup.Invoke(new ChangePanelVisibleForFinishedDelegate(ChangePanelVisibleForFinished), new object[] { });
            }
            else
            {
                this.pnlSetup.Visible = false;
                this.pnlFinish.Visible = true;
            }
        }
    }
}
