# golang-with-design-patterns: 使用Go语言实现23种设计模式

众所周知，懂得设计模式是每一个有志于从事服务端开发的萌新所必须掌握的知识；并且，设计模式在开发中也是很常用的（例如我之前曾经写过的一篇文章：《使用责任链模式解决多种用户同一个入口登录的问题》）。但是现在虽然市面上讲设计模式的文章、视频、书籍非常多，但是用Go语言描述的却屈指可数，且由于Go语言自身的一些特点，导致很多书中的代码都需要经过大刀阔斧的改造才能运用到Go语言中。那么趁着三刷《设计模式之禅》这个机会，我决定将《设计模式之禅》中的代码使用Go语言重写，对于我在项目中使用过的一些设计模式我也会尝试使用项目中的场景来进行解释，以此达到既学习设计模式，同时又学习Go语言语法与特性的目的。

如果对你有帮助的话欢迎帮忙点个Star和Fork，会给我很大很大的鼓励！感谢你的支持~如果在阅读的过程中发现错误或者发现更好的写法欢迎提出Issue或者联系我，我会第一时间进行改正。

参考资料：

- 《设计模式之禅（第2版）》——秦小波
- [5小时go语言了解设计模式（23个完整）](https://www.bilibili.com/video/BV1GD4y1D7D3)
- 其余的参考资料会在每篇文章的最后列出。

以下是一个施工进度表：

✅：已完成；🚧：施工中；❌：未开始。

- 单例模式✅
- 工厂模式✅
- 抽象工厂模式✅
- 模板模式🚧
- 建造者模式❌
- 代理模式❌
- 原型模式❌
- 中介者模式❌
- 命令模式❌
- 责任链模式❌
- 装饰模式❌
- 策略模式❌
- 适配器模式❌
- 迭代器模式❌
- 组合模式❌
- 观察者模式❌
- 门面模式❌
- 备忘录模式❌
- 访问者模式❌
- 状态模式❌
- 解释器模式❌
- 享元模式❌
- 桥梁模式❌

