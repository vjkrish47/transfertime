# transfertime
**Adaptive Compression Scheduling for Network Efficient Data Transfers**

* Author: Vijaya Krishna Namala
* Published In : International Journal of Innovative Research and Creative Technology (IJIRCT)
* Publication Date: Aug 2023
* E-ISSN: 2454-5988 
* Impact Factor : 9.142

**Abstract:**
Modern distributed systems transfer large volumes of data, but static approaches such as always compressed or uncompressed transfers lead to inefficient resource utilization. This work analyzes the trade off between bandwidth consumption and computational overhead during data movement. It demonstrates that compression significantly reduces transfer time by lowering payload size, especially for medium and large files. Experimental results show improved scalability and efficiency, highlighting the need for adaptive, runtime aware transfer strategies in distributed environments.

**Key Contributions**
* **Adaptive Compression Scheduling Framework:**\
Developed a runtime aware data transfer mechanism that dynamically decides whether to apply compression based on file size, network conditions, and system resource availability.
* **Optimized Bandwidth and Processing Trade Off:**\
Designed an approach that balances communication cost and computational overhead to minimize overall transfer time in distributed environments.
* **Comparative Transfer Performance Analysis:**\
Performed detailed evaluation comparing uncompressed and compressed transfers across multiple file sizes and network conditions to quantify efficiency gains.
* **Simulation Driven Implementation Model:**\
Implemented a Go based simulation framework to analyze concurrent transfers, network contention, and compression impact on system performance.

**Relevance & Real World Impact**
* **Reduced Data Transfer Time :**\
Adaptive compression significantly lowers transfer latency by reducing payload size, especially for medium and large files in distributed systems.
* **Efficient Bandwidth Utilization :**\
By minimizing unnecessary data transmission, the approach reduces network congestion and improves overall communication efficiency.
* **Improved System Performance :**\
Balanced use of compression avoids excessive processor overhead while maintaining faster data movement across nodes.
* **Enhanced Scalability in Distributed Environments :**\
The framework supports efficient data transfer as workload size increases, enabling stable and scalable system behavior.
* **Academic and Practical Contribution :**\
Provides a foundation for research in adaptive data transfer strategies and supports real world implementation in cloud, storage, and distributed platforms.

**Experimental Results (Summary)**:

  | File Size (MB) | Unccompressed Transfer (ms) |  Compressed Transfer (ms))| Improvement (%) |
  |----------------|-----------------------------| --------------------------| ----------------|
  | 100            | 820                         | 480                       | 41.46           |
  | 300            | 2350                        | 1320                      | 43.83           |
  | 500            | 3800                        | 2050                      | 46.05           |
  | 700            | 5200                        | 2750                      | 47.12           |
  | 900            | 6650                        | 3400                      | 48.87           |

**Citation** \
Adaptive Compression Scheduling for Network Efficient Data Transfers. \
Vijaya Krishna Namala \
International Journal of Innovative Research and Creative Technology \
E-ISSN- 2454-5988 \
**License** \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijirct.org/ \
**Author Contact** \
**LinkedIn**: linkedin.com/in/vijaya-krishna-namala-a42b2958 | **Email**: vijaya.namala@gmail.com



