# **Edge Security Challenges**


__By Kilton Hopkins, Jono Bergquist, Bernhard Ortner, Moritz Kröger, Steve Wong__

__Status:__ 
* __3/27/2019 - Initial draft__
* __5/7/2019 - First integration of contributions__
* __6/10/2019 - Second integration of contributions__
* __7/26/2019 - First published version 1.0__


## 1 Scope of this document

The purpose of this document is to describe a set of edge security challenges and concerns that the Kubernetes community feels is mostly complete. Using this format, we will avoid trying to both expose and solve security issues in the same document, which is likely to result in long timeframes for paper completion.


### 1.1 Goals



*   Identify the set of security challenges at the edge that are universal (cover roughly 80% of the total security concerns for all use cases)
*   Describe each security challenge in a manner that allows all professionals with a moderate level of technical skill to understand the issue


### 1.2 Non-goals



*   Limit the scope of any security discussions outside of this paper
*   Recommend any particular course of action regarding edge security
*   Attempt to definitively solve the security challenges listed herein


## 2 Trusting hardware

All edge computing rests upon the edge hardware. There is a trend toward increasing computing power coupled with lower energy consumption, which is enabling a rise in the count and ability of edge devices. And although security is a popular topic, there does not appear to be as strong of a trend toward enhancing hardware security.

Even when hardware security features are present, the incorporation of that security into the software layers of systems is frequently missing. Being able to trust edge hardware is the very underpinning of secure edge computing but it only has value if integrated fully.


### 2.1 Hardware root of trust is a starting point

Various methods exist for providing a hardware-based root of trust that can be relied upon for building a trusted edge computing node (e.g. [TPM](https://en.wikipedia.org/wiki/Trusted_Platform_Module), [HSM](https://en.wikipedia.org/wiki/Hardware_security_module), [RIoT security architecture](https://www.microsoft.com/en-us/research/wp-content/uploads/2016/06/RIoT20Paper-1.1-1.pdf)). Ideally the mechanism for hardware root of trust will serve as a foundational building block upon which software security layers can rely. While there is great value in end-to-end hardware verification solutions on their own, if the root of trust cannot be extended into the software infrastructure directly at the edge then a completely secure system may not be possible because of the gap. An example of this scenario would be a cloud-based hardware verification portal wherein the hardware identities are provisioned and controlled directly with no opportunity to verify the hardware in tandem with other security mechanisms.

When properly available for integration, the hardware root of trust becomes the starting identity on top of which the other software-based components can be built. It may be possible to use the hardware identity as a portion of an integrated edge node signature, but it also may just be sufficient to validate the hardware identity in the edge software and take actions upon failure. No matter how the integration is achieved, it is a starting point for the whole system, as deploying secure edge microservices to untrusted hardware defeats much of the purpose.


### 2.2 Condition of hardware

A challenge of edge hardware is to guarantee integrity of a reliable source of information about the condition of the device. Using insecure hardware condition information may lead to an increase in vulnerabilities.

 

To identify insecure devices, information about the amount of RAM, CPU, and storage available on an edge node can be stored and serve as valuable data source for purposes or orchestration and administration. Furthermore, it is also valuable for identifying potential security violations for example, unexpected increases in CPU consumption may indicate that an unauthorized process is being executed.

 

Other hardware attributes such as battery level or location (GPS coordinates) can play a role, as well. For example the battery level can be used to avoid a break in secure data streams by reallocating processing or a change of the device’s GPS coordinate changes can be used to force a shutdown of that particular software due to the fact that the device may be compromised.


### 2.3 Physical Component Access

For IoT/edge deployments, physical security can no longer make the same guarantees that data center deployments are able to. Devices live in the field and anyone with the proper motivation will find a way to access the hardware. With this in mind, it’s important to consider low-level components on a mainboard that become viable targets. SPI flash ROM, EEPROM and other non-volatile flash storage that typically hold the mainboard firmware require specialized tools to access and re-program but should still be considered vulnerable in edge scenarios with reduced physical security guarantees. Even debugging interfaces like JTAG are open to manipulation or leaking of sensitive hardware data. This validates the idea of using hardware security devices like a TPM to form the initial root of trust and validate that these mainboard components haven’t been compromised.

While TPMs provide a potential solution for mainboard components that can be accessed, there is still the concern of ports like USB, SATA, and other common buses that are not included in the original hardware design being accessed by a motivated attacker. Protecting against unauthorized (and likely unexpected) external devices accessing edge nodes via these peripheral interfaces is more important at the edge than in physically secure data centers.


### 2.4 Indication of compromise

Some edge hardware designs include intrusion detection mechanisms, often related to the opening of casing. In some cases the intrusion simply results in the edge computer refusing to boot. In other cases the intrusion results in a notification to the owner of the hardware. For a completely secure edge computing system, any indication of hardware compromise should probably result in immediate software actions being taken, as well. The challenge is to remove important information (data, software, etc.) upon detecting compromise and before it can be accessed by an unauthorized person.


### 2.5 Authenticity of hardware

Even when edge hardware contains a root of trust technology implementation, and even if that root of trust is integrated with software layers above it, there may still be a foundational breach of security if the authenticity of the hardware cannot be assured. A mainboard that masks espionage functionality as a trusted device will appear to be non-compromised but it was, in fact, designed to be untrustworthy from the start.

While these threats are typically only reserved for nation-state threat actors that have the sovereignty, influence and resources to launch an attack of this sophistication and addressing this security challenge goes far beyond the scope of technology or architecture and reaches into the areas of business operations and supply chain management, it is noted here for good reasonThe cost of replacing edge hardware, once deployed, can be much higher than replacing hardware in a centralized location. Additionally, the opportunity to discover covert edge hardware functionality may be much lower than if the hardware were present in a controlled physical location. The owner may simply never be physically near enough to the edge device during production operation in order to detect unexpected behavior.


## 3 Trusting connected devices

Edge compute nodes are frequently intended for processing data closer to the source. For many edge and IoT use cases, the source of data is not physically integrated with the edge compute hardware. Devices such as wireless sensors and IP video cameras report their data on a network shared with the edge compute node. Being able to trust the entire set of connected devices is important for mission-critical systems and should be strongly considered for almost every edge situation. Unfortunately it is not simple and many of the challenges related to this topic require different vendors working together in concert to produce scalable solutions.


### 3.1 Verifying devices and detecting corruption

It is difficult to verify the identity of a sensor or other device connected to an IoT gateway or wireless base station, yet this is necessary for creating a trusted edge data source. One problem is that device identities must be presented in a way that can be verified automatically. This problem can be addressed by introducing a digital signature. The signatures from edge compute nodes (and even sensor devices) are an effective way to verify the identity of the node. They build upon a key that is related to the hardware of the device and must be accessible to the software or firmware that produces the signature. Because this key is only intended to be used by the trusted edge compute node, activities performed with that identity are usually trusted fully.

Currently most off-the-shelf devices are not designed with a unique identity in the field. For devices that do present an identity, one must also consider the possibility of spoofing or appropriation of that identity for malicious purposes. In some cases the edge hardware is easily physically accessible, possibly even within the reach of the general public. If the trusted identity is copied by an attacker, an attacker can impersonate the edge compute node. The challenge is to ensure that the identity being trusted can be guaranteed to have integrity.

One approach to increase the integrity of a device, is to restrict the modification of data to certain formats, lengths, or volumes. This reduces the likelihood of that configuration-channels become an attack surface for various attack scenarios, but it still does not solve the issue of adversarial data that appears in the appropriate format. An example is a system configuration is overwritten and adjusts coolant device automatically in response to temperature increases of a motor. As a result, an attacker can starve the motor of coolant and possibly engender a machine failure. This issue can be addressed by adding a verification and validation step before accepting altered data.


### 3.2 Protecting data and commands

Once data enters an edge compute node, it may be stored, copied, forwarded, and used in any number of analyses. Outside of the software that receives the original data, it is difficult to guarantee the integrity as it travels onward. The same challenge exists for commands flowing to edge devices. Some commands, such as door latch actuators, should be carefully protected because spoofing or intercepting the command could result in tangible damages or losses. Mission-critical communications such as a “stop” or “kill” signal in manufacturing will likely need to be delivered with minimal delay, yet security layers used to verify the signal can introduce latency. Methods of protecting data and commands will likely involve encryption, but establishing an encrypted communication channel with a verified device identity remains an issue.


### 3.3 Device management

The concept of device management includes a very large set of activities and functionality. Everything from provisioning devices for use in a solution to monitoring battery life and updating firmware can be placed under this topic. Many device management systems accept information about the state of a device without question. In some cases, the point of ingestion for device status information is an open API available at a public IP address. For production systems, this poses a significant security threat because

 false device status information may result in physical actions. Losses in the form of time, wages, materials, and fuel are possible. But additionally the system may be distributing commands or other information to false devices which are indistinguishable from legitimate devices.


## 4 Operating system

It may be more appropriate to place some of these items into different categories but for now I have placed them in with operating system concerns. Solutions for secure operating systems exist but much like hardware root of trust implementations, they are often not integrated with greater software infrastructure and almost certainly not with microservices infrastructure at the edge.


### 4.1 BIOS and secure boot

During secure boot, the early stage drivers and boot loader are hashed by the BIOS and can be verified against an expected hash value stored in a hardware root of trust module like TPM. This is a great start and can be especially effective if a signed report of the hash comparison is sent off the machine as attestation.

Any later stage drivers and user space software will be the subject of security technologies that apply after secure boot, of course. The secure boot process accomplishes only what its name describes. But unfortunately secure boot also relies on BIOS software that executes all of the steps. Protecting the integrity of the BIOS is therefore a key part of getting an edge computing node from powering on through the part where software on top of the operating system can take over. In many cases, off-the-shelf mainboards allow a BIOS update through a set of actions that the user can take with a keyboard attached to a USB port. Single board computers and hardware designed for the edge is usually more protected, fortunately. In any case, the challenge is to protect the BIOS as thoroughly as possible because it can link the hardware root of trust and operating system into a secure chain but can be a very difficult attack point to detect otherwise.


### 4.2 Running processes and binary attestation

Once the later stage operating system drivers are loading and user space software is beginning to activate, most edge compute nodes are expected to be somewhat flexible in terms of running binaries. After all, there is little need for edge computing infrastructure if the processing taking place at the edge can be permanently planned in advance and will not change.

Because edge software is dynamic, we can monitor the binaries that get loaded at the close of the standard secure boot process and afterward but we cannot simply block non-whitelist software. An attestation method may be most appropriate here, but regardless of the technique some form of remote awareness and control of running processes is needed. The security challenge does not end with this, however, because it is likely that rogue software would falsify or cancel reporting of process monitoring as an immediate step once invoked. The challenge in this situation includes finding a way to take action immediately at the edge for the detection of unauthorized software running directly on the operating system.


### 4.4 Component Firmware Vulnerabilities

This is sometimes grouped into the same challenge as BIOS/Secure Boot concerns but the reality is this is a separate and somewhat unsolvable challenge. Many of the features that may be deployed when designing a secure system built from hardware and software components tied together by cryptography have also been implemented by the suppliers of the most fundamental components of a device’s hardware. At first blush this is fantastic that suppliers of CPUs like Intel and AMD have designed components that are secure but the reality is that their [security systems are a black box](https://boingboing.net/2016/06/15/intel-x86-processors-ship-with.html) as are any additional features designed into these systems or vulnerabilities that may exist. It doesn’t matter how secure of a system that you built on top of these core components, if they are exploitable, none of the security layers you carefully constructed on top will matter. This was a big enough concern that the NSA specifically asked those manufacturers to give the agency a way to completely [disable the systems](https://www.csoonline.com/article/3220476/researchers-say-now-you-too-can-disable-intel-me-backdoor-thanks-to-the-nsa.html).


### 4.5 Security Updates of the Operating System

Once a system is started, the necessity for a security update increases. For a secure delivery of updates to the operating system, the update mirror has to be trust and verified. This configuration will not change often, and therefore can usually be embedded into the edge OS. 

To help secure the delivery channel for the updates, https can be used instead of http. This helps avoid any alterations during the download of the software package. However, once the package is available on the system, it still must be verified, e.g. by checking the signature, before replacing the binary of the system. \



### 4.6 Audit Trail and Log Files

The operating system has to log all occurred events with the intention of facilitating an audit. These log files have to be protected in the same manner as regular data. A single event includes the following information but is not limited to:



*   Timestamp of the event
*   Process identifier
*   Access to the data/binaries
*   Configuration or source of origin
*   Event itself
*   Broad categorization of the event (if desired)


## 5 Network concerns

This section is by no means considered complete and will hopefully grow to include all known security issues related to networking at the edge. Some of the issues here have arisen simply due to the fact that edge networks are not protected in the same way as data center LANs. Some of the issues have emerged because distributed systems architectures add complexity.


### 5.1 Open ports

One method of allowing peer-to-peer connections between edge nodes is to open network ports so edge software can expose APIs. On a LAN in a factory, this should not be a problem because the network is closed and therefore standard network security measures can be relied upon. But if the edge node is in a retail store and connected by WiFi, attackers may be able to gain access to the edge node using the open ports or may launch a local denial-of-service attack. It is important to limit port-based addressing to within a hardware device at minimum and/or even down to a logical unit such as a Kubernetes Pod.


### 5.2 Fixed VPNs

To avoid exposing network ports to the public, it is common to use VPNs to interconnect sites or provide mobile edge compute nodes, such as vehicles, with a connection to backend resources. Just like fixed private keys, fixed VPNs can be a very serious security threat because a stolen VPN connection may provide direct access to private networks.


### 5.3 Network access control

Controlling the network access is a crucial aspect of protecting the connected devices, services and data against its unintentional usage. To avoid such requests, a restriction of 

these resources for example via an access control list of trusted and secured identities is a viable approach. A secure identity is a trusted person or process that facilitates

a secure communication between multiple parties.  

Once this precondition is ensured, the device or service has to properly handle the requests of the ACLs other non-trusted requests are dropped. To avoid a theft of an identity a periodic re-evaluation or change is recommended.

In addition, the connection to the network can be secured via access credentials.

They are used to prevent non-authorized connections from various services or databases but those credentials are not tied to any particular network identity in most cases. Possessing the credentials is equivalent to becoming that identity. Even without access credentials, being on the network allows an attacker to occupy the software with invalid requests in a denial-of-service attack.


### 5.4 Identity verification of control plane

It is just as important for an edge compute node to verify the identity of its control plane contacts as it is for the control plane to verify the identity of the edge nodes. A successfully impersonated control plane server has full control over the edge nodes, so steps should be taken to prevent such identity attacks.

Most systems are designed with devices and gateways deployed at the edge but continue to keep the control plane centralized. There are some scenarios where deploying the control plane at the edge is ideal but comes with significantly more risks. Control planes are made up of device databases, secret stores, and system logic that can manipulate the entire network. Running this system at the edge is an important step forward to realizing a universal compute platform but designing a system that can be secure from physical hardware access, attackers inserting themselves into the network and impersonated or compromised nodes on the network requires significantly more thought than deploying a compute platform in a secure data center.


### 5.5 Attacks of transport layer

When working with communication protocols like Zigbee, Wifi or Bluetooth it is possible to disturb or attack the communication transport layer. By using for example white noise generators the communication between edge-devices and master nodes can be complete blocked and would look similar to something like a denial of service attack. By introducing noise into a network layer some systems will automatically increase the SNR ratio which would lead to an increase in power consumption.


### 5.6 Denial-of-thing attacks

Sensors and actuators at the edge may be an attack vector but they may also be the end target of an attack, as well. By taking advantage of the less complex nature of many sensor devices, even an encrypted connection with a verified identity can be compromised. As a side note, it'd also may be more suitable to label this description as “denial-of-device attacks”.

Many wireless sensor modules use Bluetooth Low Energy, Zigbee, ZWave, or another form of low-power wireless communication. For example, let’s consider that we have a Bluetooth Low Energy (BLE) fluid level meter that signals when a given reservoir has reached capacity. BLE devices act as servers, waiting for an external device to request a connection. For this example we will assume that this particular BLE device forms a secure connection with devices that can prove authorization to connect by using an access token and a signing identity. This secure functionality is rarely seen in the field but in our thought experiment it provides us with a rather secure wireless sensor.

If the sensor is powered by a battery we may be able to bring the sensor offline by bombarding it with a series of requests to connect that fail security checks. Our goal as the attacker is not to be granted a data connection but rather to cause the device to use its wireless radio so frequently that it goes offline ahead of any maintenance schedule. Additionally, even if the device is powered by a reliable source such as a wall outlet, we may be able to occupy its limited processing capacity with the same type of bombardment and deny it the ability to form a meaningful connection with the rest of the system.

Preventing denial-of-thing attacks will probably require changes to the firmware in simple devices. Most edge use cases make use of devices from multiple vendors. To date there has been no unified approach to solving this problem.


## 6 Edge microservices

The goal of all edge infrastructure is to enable dynamic software to operate in a secure and reliable environment. Once that is achieved, the dynamic software (edge microservices) must be secured, as well.


### 6.1 Purity of images

Integrity of the edge microservice images must be verified before they are allowed to run. The microservice has to be optimized and secured in a way, that environmental footprint is locked and kept to a minimum to avoid any unintentional usage of its resources. Otherwise an edge compute node is an open platform for running arbitrary code.


### 6.2 Secure delivery of secrets

Many edge microservices need configuration, API keys, database credentials, and other information that has to be kept secret. If the microservices are built with that information embedded, then the microservice image itself is a security threat. Preferably, edge microservices will receive their sensitive information at runtime through a secure channel and only after their integrity and identity has been verified. In addition, the critical configuration information has to be exchangeable, even during the execution of the microservice to facilitate updates or renewal secure information.


### 6.3 Unauthorized microservices

Any unauthorized edge microservices should be detected and eliminated. The best approach is to prevent any microservice that is not whitelisted from ever starting, of course. Attempts to launch unauthorized microservices should be reported even if prevented successfully.

A further hurdle for access violations is, to run the microservice with the least required privileges. 


### 6.4 Controlled access to resources

Attached devices, such as serial ports and wireless radios, should only be accessible to edge microservices if specifically allowed. As more and more third party software will be operated at the edge, the opportunities for attackers to scan edge compute nodes and find vulnerabilities will increase. Volume mapping for edge microservices provides some good control over access to disk. Similar access control mechanisms need to be in place for other resources to avoid that multiple resources compete against host resources. Furthermore, the authors encourage that the consumed resources are proactively monitored to detect unusual patterns or resource bottlenecks. 


### 6.5 Guaranteed remote shutdown

Starting edge microservices remotely from a control plane is excellent but guarantees about the cessation of those microservices and the treatment of their logs and data is important for security. The deletion solution has to take care of a proper management and storage of the files or a proper closing of the log transmission to its log server. Some microservices may only be scheduled to run for short time periods because they possess sensitive information or contain precious intellectual property. Having assurance that the schedule is enforced would be a great benefit.


### 6.6 Matching microservices to edge hardware

Running a sensitive edge microservice on the wrong edge node can be a serious security threat. Most compute resources in cloud and data center environments are interchangeable because they operate on a network and precise physical location is therefore not important. Edge compute nodes, however, may not be interchangeable. They are often physically connected to specific sensors or actuators, such as a door lock, and therefore need some guarantee that the intended edge microservice is running on the designated edge node at the correct time.


### 6.7 Unauthorized outbound

Edge microservices may be authorized to run on a particular edge compute node but still do bad things once running. Attempts to protect edge data fall apart if edge microservices access the data in an authorized way but then transmit that data outside of the expected boundaries. Monitoring and preventing unwanted outbound connections is of critical importance to ensure that the edge infrastructure does not become the perfect opportunity for an attacker to move data out a side channel.


### 6.8 Distributed Container Images

Edge environments benefit greatly from having container images available nearby. Securely distributing container images to the edge requires several layers of security issues to be solved simultaneously, however. The replication of container images from the source of updates to the edge repositories may be a viable channel for an attacker to interrupt, intercept, or perform a replacement. Verifying container images at the edge means having trust in the verification process. That process most likely cannot require a connection back to cloud and data center in order to serve the edge environment properly, but that increases the importance of protecting the edge infrastructure, which may all be physically co-located and therefore easier to compromise.
