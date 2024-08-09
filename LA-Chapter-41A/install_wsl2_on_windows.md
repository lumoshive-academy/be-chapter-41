# How to Install WSL2 on Windows

## Prerequisites
Before starting the installation, ensure that you are running Windows 10, version 1903 or higher, with Build 18362 or higher. If you're using Windows 11, WSL2 is already supported.

## Step 1: Enable the Windows Subsystem for Linux (WSL)
1. Open **PowerShell** as Administrator. You can do this by right-clicking on the Start menu and selecting **Windows PowerShell (Admin)**.
2. Run the following command to enable WSL:

    ```powershell
    dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
    ```

## Step 2: Enable Virtual Machine Platform
1. In the same PowerShell window, run the following command to enable the Virtual Machine Platform:

    ```powershell
    dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
    ```

2. Restart your computer to apply the changes.

## Step 3: Install the WSL2 Linux Kernel Update Package
1. Download the latest WSL2 Linux kernel update package from [this Microsoft link](https://aka.ms/wsl2kernel).
2. Run the installer once the download is complete.

## Step 4: Set WSL2 as the Default Version
1. In the PowerShell window, set WSL2 as the default version by running:

    ```powershell
    wsl --set-default-version 2
    ```

## Step 5: Install a Linux Distribution
1. Open the **Microsoft Store**.
2. Search for your preferred Linux distribution, such as Ubuntu, Debian, or Kali Linux.
3. Click **Install** to download and install the distribution.

## Step 6: Set Up Your Linux Distribution
1. Once the installation is complete, launch the Linux distribution from the Start menu.
2. Follow the on-screen instructions to create a user account and set up the environment.

## Step 7: Verify the Installation
1. Open your Linux distribution and run the following command to verify that WSL2 is being used:

    ```bash
    wsl --list --verbose
    ```

   You should see your Linux distribution listed with the version set to `2`.

## Additional Tips
- To switch between WSL1 and WSL2 for a specific distribution, use the following command:

    ```powershell
    wsl --set-version <DistroName> 2
    ```

- Replace `<DistroName>` with the name of your installed Linux distribution.

## Conclusion
You have now successfully installed WSL2 on your Windows system. Enjoy using Linux alongside Windows!

