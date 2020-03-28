namespace web_access_local_setup
{
    partial class MainForm
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.lblLocalPort = new System.Windows.Forms.Label();
            this.lblAuthUrl = new System.Windows.Forms.Label();
            this.lblDirPath = new System.Windows.Forms.Label();
            this.txtLocalPort = new System.Windows.Forms.TextBox();
            this.txtAuthUrl = new System.Windows.Forms.TextBox();
            this.dlgDirPath = new System.Windows.Forms.FolderBrowserDialog();
            this.txtDirPath = new System.Windows.Forms.TextBox();
            this.btnSelectDirPath = new System.Windows.Forms.Button();
            this.btnSetup = new System.Windows.Forms.Button();
            this.lblStatus = new System.Windows.Forms.Label();
            this.pnlSetup = new System.Windows.Forms.Panel();
            this.pnlFinish = new System.Windows.Forms.Panel();
            this.lblFinished = new System.Windows.Forms.Label();
            this.btnClose = new System.Windows.Forms.Button();
            this.pnlFinish.SuspendLayout();
            this.SuspendLayout();
            // 
            // lblLocalPort
            // 
            this.lblLocalPort.AutoSize = true;
            this.lblLocalPort.Location = new System.Drawing.Point(80, 35);
            this.lblLocalPort.Name = "lblLocalPort";
            this.lblLocalPort.Size = new System.Drawing.Size(77, 12);
            this.lblLocalPort.TabIndex = 0;
            this.lblLocalPort.Text = "local port: ";
            // 
            // lblAuthUrl
            // 
            this.lblAuthUrl.AutoSize = true;
            this.lblAuthUrl.Location = new System.Drawing.Point(38, 78);
            this.lblAuthUrl.Name = "lblAuthUrl";
            this.lblAuthUrl.Size = new System.Drawing.Size(119, 12);
            this.lblAuthUrl.TabIndex = 1;
            this.lblAuthUrl.Text = "authentication url:";
            // 
            // lblDirPath
            // 
            this.lblDirPath.AutoSize = true;
            this.lblDirPath.Location = new System.Drawing.Point(56, 122);
            this.lblDirPath.Name = "lblDirPath";
            this.lblDirPath.Size = new System.Drawing.Size(101, 12);
            this.lblDirPath.TabIndex = 2;
            this.lblDirPath.Text = "setup directory:";
            // 
            // txtLocalPort
            // 
            this.txtLocalPort.Location = new System.Drawing.Point(176, 32);
            this.txtLocalPort.Name = "txtLocalPort";
            this.txtLocalPort.Size = new System.Drawing.Size(100, 21);
            this.txtLocalPort.TabIndex = 3;
            // 
            // txtAuthUrl
            // 
            this.txtAuthUrl.Location = new System.Drawing.Point(176, 75);
            this.txtAuthUrl.Name = "txtAuthUrl";
            this.txtAuthUrl.Size = new System.Drawing.Size(341, 21);
            this.txtAuthUrl.TabIndex = 4;
            // 
            // txtDirPath
            // 
            this.txtDirPath.Location = new System.Drawing.Point(176, 119);
            this.txtDirPath.Name = "txtDirPath";
            this.txtDirPath.ReadOnly = true;
            this.txtDirPath.Size = new System.Drawing.Size(259, 21);
            this.txtDirPath.TabIndex = 5;
            // 
            // btnSelectDirPath
            // 
            this.btnSelectDirPath.Location = new System.Drawing.Point(442, 116);
            this.btnSelectDirPath.Name = "btnSelectDirPath";
            this.btnSelectDirPath.Size = new System.Drawing.Size(75, 23);
            this.btnSelectDirPath.TabIndex = 6;
            this.btnSelectDirPath.Text = "select";
            this.btnSelectDirPath.UseVisualStyleBackColor = true;
            this.btnSelectDirPath.Click += new System.EventHandler(this.btnSelectDirPath_Click);
            // 
            // btnSetup
            // 
            this.btnSetup.Location = new System.Drawing.Point(213, 174);
            this.btnSetup.Name = "btnSetup";
            this.btnSetup.Size = new System.Drawing.Size(134, 23);
            this.btnSetup.TabIndex = 7;
            this.btnSetup.Text = "setup";
            this.btnSetup.UseVisualStyleBackColor = true;
            this.btnSetup.Click += new System.EventHandler(this.btnSetup_Click);
            // 
            // lblStatus
            // 
            this.lblStatus.AutoSize = true;
            this.lblStatus.Location = new System.Drawing.Point(13, 228);
            this.lblStatus.Name = "lblStatus";
            this.lblStatus.Size = new System.Drawing.Size(35, 12);
            this.lblStatus.TabIndex = 8;
            this.lblStatus.Text = "ready";
            // 
            // pnlSetup
            // 
            this.pnlSetup.Location = new System.Drawing.Point(15, 13);
            this.pnlSetup.Name = "pnlSetup";
            this.pnlSetup.Size = new System.Drawing.Size(544, 212);
            this.pnlSetup.TabIndex = 9;
            // 
            // pnlFinish
            // 
            this.pnlFinish.Controls.Add(this.btnClose);
            this.pnlFinish.Controls.Add(this.lblFinished);
            this.pnlFinish.Location = new System.Drawing.Point(176, 12);
            this.pnlFinish.Name = "pnlFinish";
            this.pnlFinish.Size = new System.Drawing.Size(237, 213);
            this.pnlFinish.TabIndex = 0;
            // 
            // lblFinished
            // 
            this.lblFinished.AutoSize = true;
            this.lblFinished.Font = new System.Drawing.Font("宋体", 26.25F, System.Drawing.FontStyle.Bold, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.lblFinished.Location = new System.Drawing.Point(47, 49);
            this.lblFinished.Name = "lblFinished";
            this.lblFinished.Size = new System.Drawing.Size(148, 35);
            this.lblFinished.TabIndex = 0;
            this.lblFinished.Text = "Finish.";
            // 
            // btnClose
            // 
            this.btnClose.Location = new System.Drawing.Point(71, 162);
            this.btnClose.Name = "btnClose";
            this.btnClose.Size = new System.Drawing.Size(75, 23);
            this.btnClose.TabIndex = 1;
            this.btnClose.Text = "close";
            this.btnClose.UseVisualStyleBackColor = true;
            this.btnClose.Click += new System.EventHandler(this.btnClose_Click);
            // 
            // MainForm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(571, 254);
            this.Controls.Add(this.pnlFinish);
            this.Controls.Add(this.lblStatus);
            this.Controls.Add(this.btnSetup);
            this.Controls.Add(this.btnSelectDirPath);
            this.Controls.Add(this.txtDirPath);
            this.Controls.Add(this.txtAuthUrl);
            this.Controls.Add(this.txtLocalPort);
            this.Controls.Add(this.lblDirPath);
            this.Controls.Add(this.lblAuthUrl);
            this.Controls.Add(this.lblLocalPort);
            this.Controls.Add(this.pnlSetup);
            this.MaximizeBox = false;
            this.Name = "MainForm";
            this.ShowIcon = false;
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
            this.Text = "web_access_local";
            this.Load += new System.EventHandler(this.MainForm_Load);
            this.pnlFinish.ResumeLayout(false);
            this.pnlFinish.PerformLayout();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label lblLocalPort;
        private System.Windows.Forms.Label lblAuthUrl;
        private System.Windows.Forms.Label lblDirPath;
        private System.Windows.Forms.TextBox txtLocalPort;
        private System.Windows.Forms.TextBox txtAuthUrl;
        private System.Windows.Forms.FolderBrowserDialog dlgDirPath;
        private System.Windows.Forms.TextBox txtDirPath;
        private System.Windows.Forms.Button btnSelectDirPath;
        private System.Windows.Forms.Button btnSetup;
        private System.Windows.Forms.Label lblStatus;
        private System.Windows.Forms.Panel pnlSetup;
        private System.Windows.Forms.Panel pnlFinish;
        private System.Windows.Forms.Button btnClose;
        private System.Windows.Forms.Label lblFinished;
    }
}